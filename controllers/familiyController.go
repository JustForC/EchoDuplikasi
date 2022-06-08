package controllers

import (
	"Kynesia/ent"
	"Kynesia/ent/family"
	"Kynesia/ent/register"
	"Kynesia/helpers"
	"Kynesia/requests"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type familyController struct {
	db *ent.Client
}

func NewFamilyController(db *ent.Client) *familyController {
	return &familyController{db}
}

func (famController *familyController) Create(c echo.Context) error {
	ctx := context.Background()
	req := new(requests.FamilyRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	if req.Job == "" {
		req.Job = "-"
	}

	// Take register cookie for id register scholarship
	scholarship, err := c.Cookie("Register")

	// Check if register cookie exist or not
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Select scholarship first!",
		})
	}

	// Convert id register scholarship to int
	scholarshipID, _ := strconv.Atoi(scholarship.Value)

	family := famController.db.Family.Create().
		SetStatus(req.Status).
		SetName(req.Name).
		SetGender(req.Gender).
		SetBirthplace(req.Birthplace).
		SetBirthdate(helpers.ConvertTime(req.Birthdate)).
		SetEducation(req.Education).
		SetJob(req.Job).
		AddRegisterIDs(scholarshipID).SaveX(ctx)

	return c.JSON(http.StatusOK, family)
}

func (famController *familyController) ReadAll(c echo.Context) error {
	ctx := context.Background()

	// Take register cookie for id register scholarship
	scholarship, err := c.Cookie("Register")

	// Check if register cookie exist or not
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Select scholarship first!",
		})
	}

	// Convert id register scholarship to int
	scholarshipID, _ := strconv.Atoi(scholarship.Value)

	families := famController.db.Family.Query().Where(family.HasRegisterWith(register.ID(scholarshipID))).AllX(ctx)

	return c.JSON(http.StatusOK, families)
}

func (famController *familyController) ReadByID(c echo.Context) error {
	ctx := context.Background()

	id := helpers.ConvertId(c.Param("id"))

	if check := famController.db.Family.Query().Where(family.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Family does not exist!",
		})
	}

	family := famController.db.Family.Query().Where(family.ID(id)).FirstX(ctx)

	return c.JSON(http.StatusOK, family)
}

func (famController *familyController) Update(c echo.Context) error {
	return nil
}

func (famController *familyController) Delete(c echo.Context) error {
	ctx := context.Background()

	id, _ := strconv.Atoi(c.Param("id"))

	if check := famController.db.Family.Query().Where(family.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Family does not exist!",
		})
	}

	famController.db.Family.DeleteOneID(id).ExecX(ctx)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Family deleted!",
	})
}

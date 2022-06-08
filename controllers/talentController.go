package controllers

import (
	"Kynesia/ent"
	"Kynesia/ent/register"
	"Kynesia/ent/talent"
	"Kynesia/helpers"
	"Kynesia/requests"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type talentController struct {
	db *ent.Client
}

func NewTalentController(db *ent.Client) *talentController {
	return &talentController{db}
}

func (talController *talentController) Create(c echo.Context) error {
	ctx := context.Background()
	req := new(requests.TalentRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return err
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

	talent := talController.db.Talent.Create().SetName(req.Name).AddRegisterIDs(scholarshipID).SaveX(ctx)

	return c.JSON(http.StatusOK, talent)
}

func (talController *talentController) ReadAll(c echo.Context) error {
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

	talents := talController.db.Talent.Query().Where(talent.HasRegisterWith(register.ID(scholarshipID))).AllX(ctx)

	return c.JSON(http.StatusOK, talents)
}

func (talController *talentController) ReadByID(c echo.Context) error {
	ctx := context.Background()

	id := helpers.ConvertId(c.Param("id"))

	if check := talController.db.Talent.Query().Where(talent.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Talent does not exist!",
		})
	}

	talent := talController.db.Talent.Query().Where(talent.ID(id)).FirstX(ctx)

	return c.JSON(http.StatusOK, talent)
}

func (talController *talentController) Update(c echo.Context) error {
	return nil
}

func (talController *talentController) Delete(c echo.Context) error {
	ctx := context.Background()

	id := helpers.ConvertId(c.Param("id"))

	if check := talController.db.Talent.Query().Where(talent.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Talent does not exist!",
		})
	}

	talController.db.Talent.DeleteOneID(id).ExecX(ctx)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Talent deleted!",
	})
}

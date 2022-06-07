package controllers

import (
	"Kynesia/ent"
	"Kynesia/ent/education"
	"Kynesia/requests"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type educationController struct {
	db *ent.Client
}

func NewEducationController(db *ent.Client) *educationController {
	return &educationController{db}
}

func (eduController *educationController) Create(c echo.Context) error {
	ctx := context.Background()
	req := new(requests.EducationRequest)

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

	education := eduController.db.Education.Create().
		SetGrade(req.Grade).
		SetName(req.Name).
		SetCity(req.City).
		SetProvince(req.Province).
		SetEnter(req.Enter).
		SetGraduate(req.Graduate).
		AddRegisterIDs(scholarshipID).SaveX(ctx)

	return c.JSON(http.StatusOK, education)
}

func (eduController *educationController) ReadAll(c echo.Context) error {
	ctx := context.Background()

	educations := eduController.db.Education.Query().AllX(ctx)

	return c.JSON(http.StatusOK, educations)
}

func (eduController *educationController) ReadByID(c echo.Context) error {
	ctx := context.Background()

	id, _ := strconv.Atoi(c.Param("id"))

	if check := eduController.db.Education.Query().Where(education.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Education does not exist!",
		})
	}

	education := eduController.db.Education.Query().Where(education.ID(id)).FirstX(ctx)

	return c.JSON(http.StatusOK, education)
}

func (eduController *educationController) Update(c echo.Context) error {
	return nil
}

func (eduController *educationController) Delete(c echo.Context) error {
	ctx := context.Background()

	id, _ := strconv.Atoi(c.Param("id"))

	if check := eduController.db.Education.Query().Where(education.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Education does not exist!",
		})
	}

	eduController.db.Education.DeleteOneID(id).ExecX(ctx)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Education deleted!",
	})
}

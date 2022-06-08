package controllers

import (
	"Kynesia/ent"
	"Kynesia/ent/register"
	"Kynesia/ent/training"
	"Kynesia/helpers"
	"Kynesia/requests"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type trainingController struct {
	db *ent.Client
}

func NewTrainingController(db *ent.Client) *trainingController {
	return &trainingController{db}
}

func (trainController *trainingController) Create(c echo.Context) error {
	ctx := context.Background()
	req := new(requests.TrainingRequest)

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

	training := trainController.db.Training.Create().
		SetName(req.Name).
		SetPeriod(req.Period).
		SetOrganizer(req.Organizer).
		SetYear(req.Year).
		SetCertificate(req.Certificate).
		AddRegisterIDs(scholarshipID).SaveX(ctx)

	return c.JSON(http.StatusOK, training)
}

func (trainController *trainingController) ReadAll(c echo.Context) error {
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

	trainings := trainController.db.Training.Query().Where(training.HasRegisterWith(register.ID(scholarshipID))).AllX(ctx)

	return c.JSON(http.StatusOK, trainings)
}

func (trainController *trainingController) ReadByID(c echo.Context) error {
	ctx := context.Background()

	id := helpers.ConvertId(c.Param("id"))

	if check := trainController.db.Training.Query().Where(training.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Training does not exist!",
		})
	}

	training := trainController.db.Training.Query().Where(training.ID(id)).FirstX(ctx)

	return c.JSON(http.StatusOK, training)
}

func (trainController *trainingController) Update(c echo.Context) error {
	return nil
}

func (trainController *trainingController) Delete(c echo.Context) error {
	ctx := context.Background()

	id := helpers.ConvertId(c.Param("id"))

	if check := trainController.db.Training.Query().Where(training.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Training does not exist!",
		})
	}

	trainController.db.Training.DeleteOneID(id).ExecX(ctx)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Training deleted!",
	})
}

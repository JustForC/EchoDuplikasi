package controllers

import (
	"Kynesia/ent"
	"Kynesia/ent/achievement"
	"Kynesia/ent/register"
	"Kynesia/requests"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type achievementController struct {
	db *ent.Client
}

func NewAchievementController(db *ent.Client) *achievementController {
	return &achievementController{db}
}

func (achievController *achievementController) Create(c echo.Context) error {
	ctx := context.Background()
	req := new(requests.AchievementRequest)

	// Binding request to request model
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	// Validate request
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

	// Creating achievement
	achievement := achievController.db.Achievement.Create().
		SetName(req.Name).
		SetLevel(req.Level).
		SetOrganizer(req.Organizer).
		AddRegisterIDs(scholarshipID).SaveX(ctx)

	return c.JSON(http.StatusOK, achievement)
}

func (achievController *achievementController) ReadAll(c echo.Context) error {
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

	// Taking all achievement
	achievements := achievController.db.Achievement.Query().Where(achievement.HasRegisterWith(register.ID(scholarshipID))).AllX(ctx)

	return c.JSON(http.StatusOK, achievements)
}

func (achievController *achievementController) ReadByID(c echo.Context) error {
	ctx := context.Background()

	// Taking id from url parameter
	id, _ := strconv.Atoi(c.Param("id"))

	// Check if achievement exist or not
	if check := achievController.db.Achievement.Query().Where(achievement.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Achievement does not exist!",
		})
	}

	// Taking selected achievement
	achievement := achievController.db.Achievement.Query().Where(achievement.ID(id)).FirstX(ctx)

	return c.JSON(http.StatusOK, achievement)
}

func (achievController *achievementController) Update(c echo.Context) error {
	return nil
}

func (achievController *achievementController) Delete(c echo.Context) error {
	ctx := context.Background()

	// Taking id from url parameter
	id, _ := strconv.Atoi(c.Param("id"))

	// Check if achievement exist or not
	if check := achievController.db.Achievement.Query().Where(achievement.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Achievement does not exist!",
		})
	}

	// Deleting selected achievement
	achievController.db.Achievement.DeleteOneID(id).ExecX(ctx)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Achievement deleted!",
	})
}

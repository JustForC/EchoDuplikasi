package controllers

import (
	"Kynesia/ent"
	"Kynesia/ent/language"
	"Kynesia/ent/register"
	"Kynesia/helpers"
	"Kynesia/requests"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type languageController struct {
	db *ent.Client
}

func NewLanguageController(db *ent.Client) *languageController {
	return &languageController{db}
}

func (langController *languageController) Create(c echo.Context) error {
	ctx := context.Background()
	req := new(requests.LanguageRequest)

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

	language := langController.db.Language.Create().
		SetName(req.Name).
		SetListen(req.Listen).
		SetRead(req.Read).
		SetTalk(req.Talk).
		SetWrite(req.Write).
		AddRegisterIDs(scholarshipID).SaveX(ctx)

	return c.JSON(http.StatusOK, language)
}

func (langController *languageController) ReadAll(c echo.Context) error {
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

	languages := langController.db.Language.Query().Where(language.HasRegisterWith(register.ID(scholarshipID))).AllX(ctx)

	return c.JSON(http.StatusOK, languages)
}

func (langController *languageController) ReadByID(c echo.Context) error {
	ctx := context.Background()

	id := helpers.ConvertId(c.Param("id"))

	if check := langController.db.Language.Query().Where(language.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Language does not exist!",
		})
	}

	language := langController.db.Language.Query().Where(language.ID(id)).FirstX(ctx)

	return c.JSON(http.StatusOK, language)
}

func (langController *languageController) Update(c echo.Context) error {
	return nil
}

func (langController *languageController) Delete(c echo.Context) error {
	ctx := context.Background()

	id := helpers.ConvertId(c.Param("id"))

	if check := langController.db.Language.Query().Where(language.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Language does not exist!",
		})
	}

	langController.db.Language.DeleteOneID(id).ExecX(ctx)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Language deleted!",
	})
}

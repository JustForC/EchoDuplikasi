package controllers

import (
	"Kynesia/ent"
	"Kynesia/ent/register"
	"Kynesia/ent/socialmedia"
	"Kynesia/helpers"
	"Kynesia/requests"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type socialMediaController struct {
	db *ent.Client
}

func NewSocialMediaRequest(db *ent.Client) *socialMediaController {
	return &socialMediaController{db}
}

func (socialController *socialMediaController) Create(c echo.Context) error {
	ctx := context.Background()
	req := new(requests.SocialMediaRequest)

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

	socialMedia := socialController.db.SocialMedia.Create().
		SetFacebook(req.Facebook).
		SetInstagram(req.Instagram).
		SetTiktok(req.Tiktok).
		SetTwitter(req.Twitter).
		AddRegisterIDs(scholarshipID).SaveX(ctx)

	return c.JSON(http.StatusOK, socialMedia)
}

func (socialController *socialMediaController) ReadAll(c echo.Context) error {
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

	socialMedia := socialController.db.SocialMedia.Query().Where(socialmedia.HasRegisterWith(register.ID(scholarshipID))).AllX(ctx)

	return c.JSON(http.StatusOK, socialMedia)
}

func (socialController *socialMediaController) ReadByID(c echo.Context) error {
	ctx := context.Background()

	id := helpers.ConvertId(c.Param("id"))

	if check := socialController.db.SocialMedia.Query().Where(socialmedia.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Social media does not exist!",
		})
	}

	socialMedia := socialController.db.SocialMedia.Query().Where(socialmedia.ID(id)).FirstX(ctx)

	return c.JSON(http.StatusOK, socialMedia)
}

func (socialController *socialMediaController) Update(c echo.Context) error {
	return nil
}

func (socialController *socialMediaController) Delete(c echo.Context) error {
	ctx := context.Background()

	id := helpers.ConvertId(c.Param("id"))

	if check := socialController.db.SocialMedia.Query().Where(socialmedia.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Social media does not exist!",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Social media deleted!",
	})
}

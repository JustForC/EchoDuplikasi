package controllers

import (
	"Kynesia/ent"
	"Kynesia/ent/register"
	"Kynesia/ent/scholarship"
	"Kynesia/ent/user"
	"context"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type registerController struct {
	db *ent.Client
}

func NewRegisterController(db *ent.Client) *registerController {
	return &registerController{db}
}

func (regisController *registerController) RegisterScholarship(c echo.Context) error {
	ctx := context.Background()

	use := c.Get("user").(*jwt.Token)
	claims := use.Claims.(*CustomClaim)

	id, _ := strconv.Atoi(c.Param("scholarshipID"))

	if check := regisController.db.Register.Query().Where(register.HasScholarshipWith(scholarship.ID(id))).Where(register.HasUserWith(user.ID(claims.ID))).ExistX(ctx); check == false {
		regisController.db.Register.Create().SetStatusOne(0).AddUserIDs(claims.ID).AddScholarshipIDs(id).SaveX(ctx)

		cookieScholarship := new(http.Cookie)

		cookieScholarship.Name = "Scholarship"
		cookieScholarship.Value = c.Param("scholarshipID")
		cookieScholarship.Path = "/"

		c.SetCookie(cookieScholarship)

		return c.JSON(http.StatusOK, echo.Map{
			"message": "Register success!",
		})
	}

	cookieScholarship := new(http.Cookie)

	cookieScholarship.Name = "Scholarship"
	cookieScholarship.Value = c.Param("scholarshipID")
	cookieScholarship.Path = "/"

	c.SetCookie(cookieScholarship)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Already registered!",
	})
}

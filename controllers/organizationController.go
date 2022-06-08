package controllers

import (
	"Kynesia/ent"
	"Kynesia/ent/organization"
	"Kynesia/helpers"
	"Kynesia/requests"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type organizationController struct {
	db *ent.Client
}

func NewOrganizationController(db *ent.Client) *organizationController {
	return &organizationController{db}
}

func (orgController *organizationController) Create(c echo.Context) error {
	ctx := context.Background()
	req := new(requests.OrganizationRequest)

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

	organization := orgController.db.Organization.Create().
		SetName(req.Name).
		SetDetail(req.Detail).
		SetPeriod(req.Period).
		SetPosition(req.Position).
		AddRegisterIDs(scholarshipID).SaveX(ctx)

	return c.JSON(http.StatusOK, organization)
}

func (orgController *organizationController) ReadAll(c echo.Context) error {
	ctx := context.Background()

	organizations := orgController.db.Organization.Query().AllX(ctx)

	return c.JSON(http.StatusOK, organizations)
}

func (orgController *organizationController) ReadByID(c echo.Context) error {
	ctx := context.Background()

	id := helpers.ConvertId(c.Param("id"))

	if check := orgController.db.Organization.Query().Where(organization.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Organizations does not exist!",
		})
	}

	organization := orgController.db.Organization.Query().Where(organization.ID(id)).FirstX(ctx)

	return c.JSON(http.StatusOK, organization)
}

func (orgController *organizationController) Update(c echo.Context) error {
	return nil
}

func (orgController *organizationController) Delete(c echo.Context) error {
	ctx := context.Background()

	id := helpers.ConvertId(c.Param("id"))

	if check := orgController.db.Organization.Query().Where(organization.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Organization does not exist!",
		})
	}

	orgController.db.Organization.DeleteOneID(id).ExecX(ctx)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Organization deleted!",
	})
}

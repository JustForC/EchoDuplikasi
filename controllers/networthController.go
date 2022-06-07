package controllers

import (
	"Kynesia/ent"
	"Kynesia/ent/networth"
	"Kynesia/helpers"
	"Kynesia/requests"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type networthController struct {
	db *ent.Client
}

func NewNetworthController(db *ent.Client) *networthController {
	return &networthController{db}
}

func (netController *networthController) Create(c echo.Context) error {
	ctx := context.Background()
	req := new(requests.NetworthRequest)

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

	networth := netController.db.Networth.Create().SetValue(req.Value).AddRegisterIDs(scholarshipID).SaveX(ctx)

	return c.JSON(http.StatusOK, networth)
}

func (netController *networthController) ReadAll(c echo.Context) error {
	ctx := context.Background()

	networth := netController.db.Networth.Query().AllX(ctx)

	return c.JSON(http.StatusOK, networth)
}

func (netController *networthController) ReadByID(c echo.Context) error {
	ctx := context.Background()

	id := helpers.ConvertId(c.Param("id"))

	if check := netController.db.Networth.Query().Where(networth.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Networth does not exist!",
		})
	}

	networth := netController.db.Networth.Query().Where(networth.ID(id)).FirstX(ctx)

	return c.JSON(http.StatusOK, networth)
}

func (netController *networthController) Update(c echo.Context) error {
	return nil
}

func (netController *networthController) Delete(c echo.Context) error {
	ctx := context.Background()

	id := helpers.ConvertId(c.Param("id"))

	if check := netController.db.Networth.Query().Where(networth.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Networth does not exist!",
		})
	}

	netController.db.Networth.DeleteOneID(id).ExecX(ctx)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Networth deleted!",
	})
}

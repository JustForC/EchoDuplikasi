package controllers

import (
	"Kynesia/ent"
	"Kynesia/ent/scholarship"
	"Kynesia/helpers"
	"Kynesia/requests"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type scholarshipController struct {
	db *ent.Client
}

func NewScholarshipController(db *ent.Client) *scholarshipController {
	return &scholarshipController{db}
}

func (scholarController *scholarshipController) Create(c echo.Context) error {
	ctx := context.Background()
	req := new(requests.ScholarshipRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	scholarship := scholarController.db.Scholarship.Create().SetName(req.Name).SetStartStepOne(helpers.ConvertTime(req.StartStepOne)).SetEndStepOne(helpers.ConvertTime(req.EndStepOne)).SetAnnounceStepOne(helpers.ConvertTime(req.AnnounceStepOne)).SetStartStepTwo(helpers.ConvertTime(req.StartStepTwo)).SetEndStepTwo(helpers.ConvertTime(req.EndStepTwo)).SetAnnounceStepTwo(helpers.ConvertTime(req.AnnounceStepTwo)).SetOnlineTest(req.OnlineTest).SetStatus(0).SaveX(ctx)

	return c.JSON(http.StatusOK, scholarship)
}

func (scholarController *scholarshipController) ReadAll(c echo.Context) error {
	ctx := context.Background()

	scholarship := scholarController.db.Scholarship.Query().AllX(ctx)

	return c.JSON(http.StatusOK, scholarship)
}

func (scholarController *scholarshipController) ReadByID(c echo.Context) error {
	ctx := context.Background()

	id, _ := strconv.Atoi(c.Param("id"))

	if check := scholarController.db.Scholarship.Query().Where(scholarship.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Scholarship does not exist!",
		})
	}

	scholarship := scholarController.db.Scholarship.Query().Where(scholarship.ID(id)).FirstX(ctx)

	return c.JSON(http.StatusOK, scholarship)
}

func (scholarController *scholarshipController) Update(c echo.Context) error {
	ctx := context.Background()

	id, _ := strconv.Atoi(c.Param("id"))

	req := new(requests.ScholarshipRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if check := scholarController.db.Scholarship.Query().Where(scholarship.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Scholarship does not exist!",
		})
	}

	scholarship := scholarController.db.Scholarship.Query().Where(scholarship.ID(id)).FirstX(ctx)

	if req.Name == "" {
		req.Name = scholarship.Name
	}
	if req.StartStepOne == "" {
		req.StartStepOne = scholarship.StartStepOne.String()[0:10]
	}
	if req.StartStepTwo == "" {
		req.StartStepTwo = scholarship.StartStepTwo.String()[0:10]
	}
	if req.EndStepOne == "" {
		req.EndStepOne = scholarship.EndStepOne.String()[0:10]
	}
	if req.EndStepTwo == "" {
		req.EndStepTwo = scholarship.EndStepTwo.String()[0:10]
	}
	if req.AnnounceStepOne == "" {
		req.AnnounceStepOne = scholarship.AnnounceStepOne.String()[0:10]
	}
	if req.AnnounceStepTwo == "" {
		req.AnnounceStepTwo = scholarship.AnnounceStepTwo.String()[0:10]
	}
	if req.OnlineTest == "" {
		req.OnlineTest = scholarship.OnlineTest
	}

	scholarController.db.Scholarship.UpdateOneID(id).SetName(req.Name).SetStartStepOne(helpers.ConvertTime(req.StartStepOne)).SetEndStepOne(helpers.ConvertTime(req.EndStepOne)).SetAnnounceStepOne(helpers.ConvertTime(req.AnnounceStepOne)).SetStartStepTwo(helpers.ConvertTime(req.StartStepTwo)).SetEndStepTwo(helpers.ConvertTime(req.EndStepTwo)).SetAnnounceStepTwo(helpers.ConvertTime(req.AnnounceStepTwo)).SetOnlineTest(req.OnlineTest).SetStatus(0).SaveX(ctx)

	return c.JSON(http.StatusOK, req)
}

func (scholarController *scholarshipController) Delete(c echo.Context) error {
	ctx := context.Background()

	id, _ := strconv.Atoi(c.Param("id"))

	if check := scholarController.db.Scholarship.Query().Where(scholarship.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Scholarship does not exist!",
		})
	}

	scholarController.db.Scholarship.DeleteOneID(id).ExecX(ctx)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Data berhasil dihapus!",
	})
}

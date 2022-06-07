package controllers

import (
	"Kynesia/ent"
	"Kynesia/ent/biodata"
	"Kynesia/helpers"
	"Kynesia/requests"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type biodataController struct {
	db *ent.Client
}

func NewBiodataController(db *ent.Client) *biodataController {
	return &biodataController{db}
}

func (bioController *biodataController) Create(c echo.Context) error {
	ctx := context.Background()
	req := new(requests.BiodataRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if req.Same_Address == true {
		req.Living_Address = req.Address
		req.Living_City = req.City
		req.Living_District = req.District
		req.Living_Post_Code = req.Post_Code
		req.Living_Province = req.Province
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

	biodata := bioController.db.Biodata.Create().
		SetName(req.Name).
		SetNickname(req.Nickname).
		SetGender(req.Gender).
		SetBirthplace(req.Birthplace).
		SetBirthdate(helpers.ConvertTime(req.Birthdate)).
		SetTelephone(req.Telephone).
		SetEmail(req.Email).
		SetIdType(req.Id_Type).
		SetIdNumber(req.Id_Number).
		SetAddressID(req.Address).
		SetAddressLiving(req.Living_Address).
		SetPostCodeID(req.Post_Code).
		SetPostCodeLiving(req.Living_Post_Code).
		SetDistrictID(req.District).
		SetDistrictLiving(req.Living_District).
		SetCityID(req.City).
		SetCityLiving(req.Living_City).
		SetProvinceID(req.Province).
		SetProvinceLiving(req.Living_Province).
		SetEntrance(req.Entrance).
		SetEntranceNumber(req.Entrance_Number).
		SetUniversity(req.University).
		SetMajor(req.Major).
		AddRegisterIDs(scholarshipID).SaveX(ctx)

	return c.JSON(http.StatusOK, biodata)
}

func (bioController *biodataController) ReadAll(c echo.Context) error {
	ctx := context.Background()

	biodata := bioController.db.Biodata.Query().AllX(ctx)

	return c.JSON(http.StatusOK, biodata)
}

func (bioController *biodataController) ReadByID(c echo.Context) error {
	ctx := context.Background()

	id, _ := strconv.Atoi(c.Param("id"))

	if check := bioController.db.Biodata.Query().Where(biodata.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Biodata does not exist!",
		})
	}

	biodata := bioController.db.Biodata.Query().Where(biodata.ID(id)).FirstX(ctx)

	return c.JSON(http.StatusOK, biodata)
}

func (bioController *biodataController) Update(c echo.Context) error {
	return nil
}

func (bioController *biodataController) Delete(c echo.Context) error {
	ctx := context.Background()

	id, _ := strconv.Atoi(c.Param("id"))

	if check := bioController.db.Biodata.Query().Where(biodata.ID(id)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Biodata does not exist!",
		})
	}

	bioController.db.Biodata.DeleteOneID(id).ExecX(ctx)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Biodata deleted!",
	})
}

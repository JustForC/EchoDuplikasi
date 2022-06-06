package routes

import (
	"Kynesia/configs"
	"Kynesia/controllers"
	"Kynesia/validations"
	"context"
	"log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	// Mengganti validator bawaan echo dengan custom validator menggunakan playground
	e.Validator = &validations.CustomValidation{Validator: validator.New()}
	middleware.ErrJWTMissing.Message = echo.Map{
		"message": "Unauthorized!",
	}
	// Connect dengan database
	db := configs.Connect()

	// Initiate Controller
	authCont := controllers.NewAuthenticationController(db)
	scholarCont := controllers.NewScholarshipController(db)

	// Buat Context kosong
	ctx := context.Background()

	// Migrate schema table ke dalam database
	if err := db.Schema.Create(ctx); err != nil {
		log.Fatalf("Failed creating schema resources: %v", err.Error())
	}

	// Routing

	// Grouping Route Authentication
	auth := e.Group("/auth")
	auth.POST("/register", authCont.Register)
	auth.POST("/login", authCont.Login)
	auth.POST("/logout", authCont.Logout)

	// Grouping Route CRUD Scholarship
	scholarship := e.Group("/scholarship")
	// e.Use(middleware.JWTWithConfig(controllers.Config()))
	scholarship.POST("", scholarCont.Create)
	scholarship.GET("", scholarCont.ReadAll)
	scholarship.GET("/:id", scholarCont.ReadByID)
	scholarship.PUT("/:id", scholarCont.Update)
	scholarship.DELETE("/:id", scholarCont.Delete)

	return e
}

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
	regCont := controllers.NewRegisterController(db)
	achievCont := controllers.NewAchievementController(db)
	bioCont := controllers.NewBiodataController(db)
	eduCont := controllers.NewEducationController(db)
	famCont := controllers.NewFamilyController(db)

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
	scholarship.PUT("/activate/:id", scholarCont.Activate)
	scholarship.PUT("/deactivate/:id", scholarCont.Deactivate)

	register := e.Group("/register")
	register.Use(middleware.JWTWithConfig(controllers.Config()))
	register.GET("/:scholarshipID", regCont.RegisterScholarship)

	achievement := e.Group("/achievement")
	achievement.POST("", achievCont.Create)
	achievement.GET("", achievCont.ReadAll)
	achievement.GET("/:id", achievCont.ReadByID)
	achievement.PUT("/:id", achievCont.Update)
	achievement.DELETE("/:id", achievCont.Delete)

	biodata := e.Group("/biodata")
	biodata.POST("", bioCont.Create)
	biodata.GET("", bioCont.ReadAll)
	biodata.GET("/:id", bioCont.ReadByID)
	biodata.PUT("/:id", bioCont.Update)
	biodata.DELETE("/:id", bioCont.Delete)

	education := e.Group("/education")
	education.POST("", eduCont.Create)
	education.GET("", eduCont.ReadAll)
	education.GET("/:id", eduCont.ReadByID)
	education.PUT("/:id", eduCont.Update)
	education.DELETE(":/id", eduCont.Delete)

	family := e.Group("/family")
	family.POST("", famCont.Create)
	family.GET("", famCont.ReadAll)
	family.GET("/:id", famCont.ReadByID)
	family.PUT("/:id", famCont.Update)
	family.DELETE("/:id", famCont.Delete)

	return e
}

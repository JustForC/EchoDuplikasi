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

	// Custom message if not login
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
	langCont := controllers.NewLanguageController(db)
	netCont := controllers.NewNetworthController(db)
	trainCont := controllers.NewTrainingController(db)
	orgCont := controllers.NewOrganizationController(db)
	socialCont := controllers.NewSocialMediaRequest(db)
	talCont := controllers.NewTalentController(db)

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
	scholarship.Use(middleware.JWTWithConfig(controllers.Config()))
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

	language := e.Group("/language")
	language.POST("", langCont.Create)
	language.GET("", langCont.ReadAll)
	language.GET("/:id", langCont.ReadByID)
	language.PUT("/:id", langCont.Update)
	language.DELETE("/:id", langCont.Delete)

	networth := e.Group("/networth")
	networth.POST("", netCont.Create)
	networth.GET("", netCont.ReadAll)
	networth.GET("/:id", netCont.ReadByID)
	networth.PUT("/:id", netCont.Update)
	networth.DELETE("/:id", netCont.Delete)

	training := e.Group("/training")
	training.POST("", trainCont.Create)
	training.GET("", trainCont.ReadAll)
	training.GET("/:id", trainCont.ReadByID)
	training.PUT("/:id", trainCont.Update)
	training.DELETE("/:id", trainCont.Delete)

	organization := e.Group("/organization")
	organization.POST("", orgCont.Create)
	organization.GET("", orgCont.ReadAll)
	organization.GET("/:id", orgCont.ReadByID)
	organization.PUT("/:id", orgCont.Update)
	organization.DELETE("/:id", orgCont.Delete)

	social := e.Group("/social")
	social.POST("", socialCont.Create)
	social.GET("", socialCont.ReadAll)
	social.GET("/:id", socialCont.ReadByID)
	social.PUT("/:id", socialCont.Update)
	social.DELETE("/:id", socialCont.Delete)

	talent := e.Group("/talent")
	talent.POST("", talCont.Create)
	talent.GET("", talCont.ReadAll)
	talent.GET("/:id", talCont.ReadByID)
	talent.PUT("/:id", talCont.Update)
	talent.DELETE("/:id", talCont.Delete)

	return e
}

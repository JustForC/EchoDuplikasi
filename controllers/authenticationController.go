package controllers

import (
	"Kynesia/ent"
	"Kynesia/ent/user"
	"Kynesia/requests"
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaim struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	jwt.StandardClaims
}

type authenticationController struct {
	db *ent.Client
}

func NewAuthenticationController(db *ent.Client) *authenticationController {
	return &authenticationController{db}
}

func (authController *authenticationController) Register(c echo.Context) error {
	// Buat Context Kosong
	ctx := context.Background()

	// Buat struct RegisterRequest kosong
	req := new(requests.RegisterRequest)

	// Bind Request Ke Dalam Struct RegisterRequest
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	// Validating Request
	if err := c.Validate(req); err != nil {
		return err
	}

	// Check Unique
	if email := authController.db.User.Query().Where(user.Email(req.Email)).ExistX(ctx); email == true {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Email already exist!",
		})
	}

	// Encrypt password
	password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	// Create User
	authController.db.User.Create().SetEmail(req.Email).SetName(req.Name).SetPassword(string(password)).SetRole(0).ExecX(ctx)

	// Return Message
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully create an account!",
	})
}

func (authController *authenticationController) Login(c echo.Context) error {
	// Buat context kosong
	ctx := context.Background()

	// Buat struct RequestLogin Kosong
	req := new(requests.LoginRequest)

	// Binding request ke struct LoginRequest
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	// Validate Input
	if err := c.Validate(req); err != nil {
		return err
	}

	// Check Account Exist or Not
	if check := authController.db.User.Query().Where(user.Email(req.Email)).ExistX(ctx); check == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "User does not exist!",
		})
	}

	// Compare User Password With Input Password
	if check := bcrypt.CompareHashAndPassword([]byte(authController.db.User.Query().Where(user.Email(req.Email)).Select(user.FieldPassword).StringX(ctx)), []byte(req.Password)); check != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Incorrect email or password!",
		})
	}

	claims := &CustomClaim{
		authController.db.User.Query().Where(user.Email(req.Email)).Select(user.FieldName).StringX(ctx),
		authController.db.User.Query().Where(user.Email(req.Email)).Select(user.FieldID).IntX(ctx),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret, _ := godotenv.Read()

	value, _ := token.SignedString([]byte(secret["SECRET_KEY"]))

	// Make new cookie
	cookieToken := new(http.Cookie)

	// Filling cookie value and name
	cookieToken.Name = "Authentication"
	cookieToken.Value = value
	cookieToken.Path = "/"

	// Save cookie
	c.SetCookie(cookieToken)

	// Return message and JWT
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login Success!",
		"token":   value,
	})
}

func (authController *authenticationController) Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:    "Authentication",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	})

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Logout Success!",
	})
}

func Config() middleware.JWTConfig {
	secret, _ := godotenv.Read()

	return middleware.JWTConfig{
		Claims:      &CustomClaim{},
		SigningKey:  []byte(secret["SECRET_KEY"]),
		TokenLookup: "cookie:Authentication",
	}
}

package controllers

import (
	"echo_rest/helpers"
	"echo_rest/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Roles string `json:"roles"`
	jwt.StandardClaims
}

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")
	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func CheckLogin(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})

	}

	if !res {
		return echo.ErrUnauthorized
	}

	// prepare payload

	// Set custom claims
	claims := &jwtCustomClaims{
		"Jon Snow",
		"admin",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

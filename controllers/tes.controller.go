package controllers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate

func TesValidasi(c echo.Context) error {
	example := struct {
		Nama string
		Umur int
	}{}

	example.Nama = "Farid"
	fmt.Println(example)

	validate = validator.New()
	myEmail := 1
	errs := validate.Var(myEmail, "required,eq=1")
	fmt.Println(errs)
	return c.JSON(200, "OK")
}

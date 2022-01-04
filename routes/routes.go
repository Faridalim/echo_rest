package routes

import (
	"echo_rest/controllers"
	"echo_rest/helpers"
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {

	e := echo.New()

	config := middleware.JWTConfig{
		Claims:     &helpers.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	//public
	e.GET("/", hello)

	e.POST("/pegawai", controllers.StorePegawai)
	e.PUT("/pegawai", controllers.UpdatePegawai)
	e.DELETE("/pegawai", controllers.DeletePegawai)

	e.GET("/generatehash", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)
	e.GET("/tes", controllers.TesValidasi)

	r := e.Group("/restricted")

	r.Use(middleware.JWTWithConfig(config))
	r.GET("/pegawai", controllers.FetchAllPegawai)
	r.GET("/updateakses", controllers.UpdateAccess)

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic("fail")
	}
	ioutil.WriteFile("routes.json", data, 0644)

	return e

}

func hello(c echo.Context) error {
	return c.String(200, "OK")
}

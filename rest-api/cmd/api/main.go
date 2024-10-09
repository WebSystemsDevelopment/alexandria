package main

import (
	_ "github.com/WebSystemsDevelopment/alexandria/rest-api/docs" // This is for Swagger documentation
	"github.com/WebSystemsDevelopment/alexandria/rest-api/internal/routes"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Alexandria Rest API
// @version 1.0
// @description This is a Rest API for managing books in a library system.
// @host localhost:8080
// @BasePath /
func main() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

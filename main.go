// main.go

package main

import (
 	"github.com/baselrabia/book-api/models"
 	"github.com/baselrabia/book-api/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Database setup
	db := models.InitDB("books.db")
	models.Migrate(db)

  
  	// Routes
	api := e.Group("/api")
	routes.SetupRoutes(api)

	e.Logger.Fatal(e.Start(":8080"))
}

package routes

import (
	"github.com/baselrabia/book-api/handlers"
	"github.com/baselrabia/book-api/models"
	"github.com/baselrabia/book-api/repository"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(g *echo.Group) {
	repo := repository.NewGormBookRepository(models.DB)
	bookHandler := handlers.NewBookHandler(repo)

	g.GET("/books", bookHandler.GetAllBooks)
	g.POST("/books", bookHandler.CreateBook)
	g.GET("/books/:id", bookHandler.GetBook)
	g.PUT("/books/:id", bookHandler.UpdateBook)
	g.DELETE("/books/:id", bookHandler.DeleteBook)
}

package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/baselrabia/book-api/models"
	"github.com/baselrabia/book-api/repository"
	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	Repo repository.BookRepository
}

func NewBookHandler(repo repository.BookRepository) *BookHandler {
	return &BookHandler{Repo: repo}
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	book := new(models.Book)
	if err := c.Bind(book); err != nil {
		return err
	}
	// Validate the book data
	if err := validateBook(book); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.Repo.CreateBook(book); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create the book.")
	}

	return c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) GetBook(c echo.Context) error {
	id, err := getIntId(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid book ID")
	}

	book, err := h.Repo.GetBookByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Book not found")
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) GetAllBooks(c echo.Context) error {
	books, err := h.Repo.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve books.")
	}
	return c.JSON(http.StatusOK, books)
}

func (h *BookHandler) UpdateBook(c echo.Context) error {
	id, err := getIntId(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid book ID")
	}

	// Parse and validate the ID
	book, err := h.Repo.GetBookByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Book not found")
	}

	if err := c.Bind(book); err != nil {
		return err
	}

	if err := h.Repo.UpdateBook(book); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update the book.")
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	pid, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("hhhhhhhhhwwwhh", 777, pid)
	c.Logger().Info("Received DELETE request for task ID: ", pid)

	id, err := getIntId(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid book ID")
	}
	fmt.Println("hhhhhhhhhhh", id)

	if err := h.Repo.DeleteBook(id); err != nil {
		return c.JSON(http.StatusNotFound, "Book not found")
	}

	return c.NoContent(http.StatusNoContent)
}
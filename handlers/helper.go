package handlers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/baselrabia/book-api/models"
	"github.com/labstack/echo/v4"
)

func validateBook(book *models.Book) error {
	if book.Title == "" {
		return errors.New("title is required")
	}

	if book.Author == "" {
		return errors.New("author is required")
	}

	if book.Published <= 0 {
		return errors.New("published year must be greater than 0")
	}

	return nil
}

func getIntId(c echo.Context) (uint, error) {
	idStr := c.Param("id")
	fmt.Println("hhhhhhhhhhh", idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

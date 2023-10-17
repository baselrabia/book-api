package repository

import (
	"github.com/baselrabia/book-api/dto"
	"github.com/baselrabia/book-api/models"
)

type BookRepository interface {
	CreateBook(book *dto.Book) error
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id uint) (*models.Book, error)
	UpdateBook(id uint, book *dto.Book) (*models.Book, error)
	DeleteBook(id uint) error
}

package repository

import 	"github.com/baselrabia/book-api/models"


type BookRepository interface {
	CreateBook(book *models.Book) error
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id uint) (*models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id uint) error
}

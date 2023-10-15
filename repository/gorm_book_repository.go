// repository/gorm_book_repository.go

package repository

import (
	"gorm.io/gorm"
	"github.com/baselrabia/book-api/models"
)

type GormBookRepository struct {
	DB *gorm.DB
}

func NewGormBookRepository(db *gorm.DB) *GormBookRepository {
	return &GormBookRepository{DB: db}
}

func (r *GormBookRepository) CreateBook(book *models.Book) error {
	return r.DB.Create(book).Error
}

func (r *GormBookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.DB.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *GormBookRepository) GetBookByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.DB.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *GormBookRepository) UpdateBook(book *models.Book) error {
	return r.DB.Save(book).Error
}

func (r *GormBookRepository) DeleteBook(id uint) error {
	return r.DB.Delete(&models.Book{}, id).Error
}

package service

import "bookstore/models/db"
import "bookstore/models/entity"

type BookService struct{}

type CreateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (b BookService) GetAll() ([]entity.Book, error) {
	var books []entity.Book
	err := db.DB.Find(&books).Error
	return books, err
}

func (b BookService) Get(id string) (entity.Book, error) {
	var book entity.Book
	err := db.DB.Where("id = ?", id).First(&book).Error
	return book, err
}

func (b BookService) Create(req CreateBookRequest) (entity.Book, error) {
	book := entity.Book{Title: req.Title, Author: req.Author}
	err := db.DB.Create(&book).Error
	return book, err
}

func (b BookService) Update(id string, req UpdateBookRequest) (entity.Book, error) {
	book, err := b.Get(id)
	if err != nil {
		return entity.Book{}, err
	}
	updateErr := db.DB.Model(&book).Updates(req).Error
	return book, updateErr
}

func (b BookService) Delete(id string) error {
	book, err := b.Get(id)
	if err != nil {
		return err
	}
	deleteErr := db.DB.Delete(&book).Error
	return deleteErr
}

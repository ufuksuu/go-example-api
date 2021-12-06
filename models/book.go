package models

type BookData struct{}

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (b BookData) GetAll() ([]Book, error) {
	var books []Book
	err := DB.Find(&books).Error
	return books, err
}

func (b BookData) Get(id string) (Book, error) {
	var book Book
	err := DB.Where("id = ?", id).First(&book).Error
	return book, err
}

func (b BookData) Create(input CreateBookInput) (Book, error) {
	book := Book{Title: input.Title, Author: input.Author}
	err := DB.Create(&book).Error
	return book, err
}

func (b BookData) UpdateBook(book *Book, input UpdateBookInput) error {
	updateErr := DB.Model(&book).Updates(input).Error
	return updateErr
}

func (b BookData) DeleteBook(book *Book) error {
	err := DB.Delete(&book).Error
	return err
}

package controllers

import (
	"bookstore/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BooksController struct{}

var bookModel = new(models.BookData)

// GetAllBooks GET /books
// Get all books in the database
func (b BooksController) GetAllBooks(c *gin.Context) {
	books, err := bookModel.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
	return
}

// GetBook GET /books/:id
// Get a book with id
func (b BooksController) GetBook(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		book, err := bookModel.Get(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": book})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "id not found"})
	c.Abort()
	return
}

// CreateBook POST /books
// Create new book
func (b BooksController) CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	book, createError := bookModel.Create(input)
	if createError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": createError.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
	return
}

// UpdateBook PATCH /books/:id
// Update book data
func (b BooksController) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		book, err := bookModel.Get(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		var input models.UpdateBookInput
		if bindErr := c.ShouldBindJSON(&input); bindErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": bindErr.Error()})
			c.Abort()
			return
		}
		updateError := bookModel.UpdateBook(&book, input)
		if updateError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": updateError.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": book})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "id not found"})
	c.Abort()
	return
}

// DeleteBook DELETE /books/:id
// Delete a book record
func (b BooksController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		book, err := bookModel.Get(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		deleteError := bookModel.DeleteBook(&book)
		if deleteError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": deleteError.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "id not found"})
	c.Abort()
	return
}

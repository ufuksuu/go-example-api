package controllers

import (
	"bookstore/models/entity"
	"bookstore/models/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BooksController struct{}

var bookService = new(service.BookService)

// GetAllBooks GET /books
// Get all books in the database
func (b BooksController) GetAllBooks(c *gin.Context) {
	books, err := bookService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{Message: err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, entity.Response{Data: books})
	return
}

// GetBook GET /books/:id
// Get a book with id
func (b BooksController) GetBook(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		book, err := bookService.Get(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, entity.Response{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, entity.Response{Data: book})
		return
	}
	c.JSON(http.StatusBadRequest, entity.Response{Message: "id not found"})
	c.Abort()
	return
}

// CreateBook POST /books
// Create new book
func (b BooksController) CreateBook(c *gin.Context) {
	var req service.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.Response{Message: err.Error()})
		c.Abort()
		return
	}
	book, createError := bookService.Create(req)
	if createError != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{Message: createError.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, entity.Response{Data: book})
	return
}

// UpdateBook PATCH /books/:id
// Update book data
func (b BooksController) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		var req service.UpdateBookRequest
		if bindErr := c.ShouldBindJSON(&req); bindErr != nil {
			c.JSON(http.StatusBadRequest, entity.Response{Message: bindErr.Error()})
			c.Abort()
			return
		}
		book, err := bookService.Update(id, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, entity.Response{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, entity.Response{Data: book})
		return
	}
	c.JSON(http.StatusBadRequest, entity.Response{Message: "id not found"})
	c.Abort()
	return
}

// DeleteBook DELETE /books/:id
// Delete a book record
func (b BooksController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		err := bookService.Delete(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, entity.Response{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, entity.Response{Data: true})
		return
	}
	c.JSON(http.StatusBadRequest, entity.Response{Message: "id not found"})
	c.Abort()
	return
}

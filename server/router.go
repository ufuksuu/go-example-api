package server

import (
	"bookstore/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	books := new(controllers.BooksController)
	router.Group("/books")
	{
		router.GET("/books", books.GetAllBooks)
		router.GET("/books/:id", books.GetBook)
		router.POST("/books", books.CreateBook)
		router.PATCH("/books/:id", books.UpdateBook)
		router.DELETE("/books/:id", books.DeleteBook)
	}

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	return router
}

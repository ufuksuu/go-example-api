package server

import (
	"bookstore/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	bookRouter := router.Group("/books")
	bookRoutes(bookRouter)

	healthRouter := router.Group("/health")
	healthRoutes(healthRouter)

	return router
}

func bookRoutes(router *gin.RouterGroup) {
	books := new(controllers.BooksController)
	router.GET("", books.GetAllBooks)
	router.GET("/:id", books.GetBook)
	router.POST("", books.CreateBook)
	router.PATCH("/:id", books.UpdateBook)
	router.DELETE("/:id", books.DeleteBook)
}

func healthRoutes(router *gin.RouterGroup) {
	health := new(controllers.HealthController)
	router.GET("", health.Status)
}

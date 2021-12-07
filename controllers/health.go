package controllers

import (
	"bookstore/models/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct{}

// Status GET /health
// Current status of api
func (h HealthController) Status(c *gin.Context) {
	c.JSON(http.StatusOK, entity.Response{Message: "Working!", Data: true})
	return
}

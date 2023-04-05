package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"status": "accepted",
	})
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusCreated, gin.H{
		"book": id,
	})
}

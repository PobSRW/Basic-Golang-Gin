package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

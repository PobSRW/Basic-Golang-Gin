package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	id, existID := c.Get("id")
	name, existName := c.Get("name")
	message, existMsg := c.Get("message")
	isAdmin, existStatus := c.Get("is_admin")

	if !existID || !existName || !existMsg || !existStatus {
		c.JSON(500, gin.H{
			"message": "cannot get exist item",
		})
		c.Abort()
	}

	fmt.Println("ID of User is :", id)
	fmt.Println("Name of User is :", name)
	fmt.Println("Message of User is :", message)
	fmt.Println("Status of User is :", isAdmin)

}

func ShowUser(c *gin.Context) {

	id, existID := c.Get("id")
	name, existName := c.Get("name")
	message, existMsg := c.Get("message")
	isAdmin, existStatus := c.Get("is_admin")

	if !existID || !existName || !existMsg || !existStatus {
		c.JSON(500, gin.H{
			"message": "cannot get exist item",
		})
		c.Abort()
	}

	if isAdmin == true {
		c.JSON(http.StatusOK, gin.H{
			"message": "you are admin",
		})
		c.Abort()
	}

	if isAdmin == false {
		c.JSON(http.StatusOK, gin.H{
			"UserID":   id,
			"UserName": name,
			"Message":  message,
		})
	}

}

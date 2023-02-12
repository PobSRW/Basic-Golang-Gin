package middleware

import (
	"fmt"
	"gin-service/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		var input utils.Result

		input.ID = c.Query("id")
		input.Name = c.Query("name")
		input.Message = c.Query("message")

		start := time.Now()
		// runtime := time.Now().Sub(start).Seconds() * 1000

		err := c.Bind(&input)
		if err != nil {
			fmt.Println(err)
		}

		if input.ID == "" {
			fmt.Println("ID is empty")
			c.Abort() //prevents pending handlers from being called.
		}

		if input.Name == "" {
			fmt.Println("Name is empty")
			c.Abort()
		}

		if input.Message == "" {
			fmt.Println("Message is empty")
			c.Abort()
		}

		// checking admin status

		if input.ID == "1001" && input.Name == "hero" {
			c.Set("is_admin", true)
		}

		if input.ID != "1001" || input.Name != "hero" {
			c.Set("is_admin", false)
		}

		// set item for pass through api
		c.Set("id", input.ID)
		c.Set("name", input.Name)
		c.Set("message", input.Message)

		// pass item with func Next
		c.Next()

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"runtime": start,
		})
	}

}

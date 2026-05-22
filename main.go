package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Name   string    `json:"name"`
}

func getGreeting() string {
	hour := time.Now().Hour()
	switch {
	case hour >= 5 && hour < 12:
		return "Good Morning,"
	case hour >= 12 && hour < 17:
		return "Good Afternoon,"
	case hour >= 17 && hour < 21:
		return "Good Evening,"
	default:
		return "Good Night,"
	}
}

func main() {
	r := gin.Default()

	// Existing Get route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I am doing this to test trivy installation!",
		})
	})
	// New Post route
	r.POST("/greet", func(c *gin.Context) {
		var req Request

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}
		
		if req.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Name field is required",
			})
			return
		}
		
		greeting := getGreeting()

		c.JSON(http.StatusOK, gin.H{
			"message": greeting + " " + req.Name,
		})
	})



	r.Run() // listen and serve on 0.0.0.0:8080
}
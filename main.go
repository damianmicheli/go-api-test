package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":"Hello, world!",
		})
	})

	router.Run(":8081")
}
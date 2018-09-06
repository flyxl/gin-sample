package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ret": 0,
			"msg": "Welcome flyxl",
		})
	})

	router.Run(":3300")
	log.Println("server is running on port 3300")
}

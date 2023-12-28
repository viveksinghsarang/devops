// Install required package:
// go get github.com/gin-gonic/gin

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/process", func(c *gin.Context) {
		var data map[string]interface{}
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Log the request and response
		log.Printf("Received data: %v\n", data)
		c.JSON(http.StatusOK, gin.H{"message": "Request processed successfully by Go service"})
	})

	if err := router.Run(":5000"); err != nil {
		log.Fatal(err)
	}
}


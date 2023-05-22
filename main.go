package main

import (
	"GO-WEB-NATIVE/config"
	"GO-WEB-NATIVE/controllers/homecontroller"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

// func main() {
// 	config.ConnectDB()

// 	//1.homepage
// 	http.HandleFunc("/", homecontroller.Welcome)

// 	//2. Categories

//		log.Println("Server running into port 8080")
//		http.ListenAndServe(":8080", nil)
//	}
func main() {
	config.ConnectDB()
	r := gin.Default()

	// 1. Homepage route
	r.GET("/", func(c *gin.Context) {
		homecontroller.Welcome(c.Writer, c.Request)
	})

	// 3. /ping route from the first code snippet
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Println("Server running on port 8080")
	r.Run(":8080")
}

package main

import (
	"GO-WEB-NATIVE/config"
	"GO-WEB-NATIVE/controllers/homecontroller"
	"GO-WEB-NATIVE/controllers/categorycontroller"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	// 1. Homepage route
	r.GET("/", func(c *gin.Context) {
		homecontroller.Welcome(c.Writer, c.Request)
	})
	// 2. Categories route
	r.GET("/categories", func(c *gin.Context) {
		categorycontroller.Index(c.Writer, c.Request)
	})
	r.GET("/categories/add", func(c *gin.Context) {
		categorycontroller.Add(c.Writer, c.Request)
	})
	r.GET("/categories/edit", func(c *gin.Context) {
		categorycontroller.Edit(c.Writer, c.Request)
	})
	r.GET("/categories/delete", func(c *gin.Context) {
		categorycontroller.Delete(c.Writer, c.Request)
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

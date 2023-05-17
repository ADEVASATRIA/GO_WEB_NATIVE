package main

import (
	"GO-WEB-NATIVE/config"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	log.Println("Server running into port 8080")
	http.ListenAndServe(":8080", nil)
}

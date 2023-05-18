package main

import (
	"GO-WEB-NATIVE/config"
	"GO-WEB-NATIVE/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()


	//1.homepage
	http.HandleFunc("/", homecontroller.Welcome)
	log.Println("Server running into port 8080")
	http.ListenAndServe(":8080", nil)
}

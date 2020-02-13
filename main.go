package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	log.Println("Listening...")

	err := http.ListenAndServe(":"+port, http.FileServer(http.Dir("public")))

	if err != nil {
		log.Printf("Error running web server for static assets: %v", err)
	}
}

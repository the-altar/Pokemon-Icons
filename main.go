package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("Listening...")
	err := http.ListenAndServe(":3000", http.FileServer(http.Dir("public")))

	if err != nil {
		log.Printf("Error running web server for static assets: %v", err)
	}
}

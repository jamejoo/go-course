package main

import (
	"log"
	"net/http"

	"github.com/jamejoo/go-course/handlers"
)

func main() {
	const port = ":8080"

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

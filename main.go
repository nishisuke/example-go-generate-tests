package main

import (
	"log"
	"net/http"
	"sample/internal/controllers"
)

func main() {
	c := controllers.NewUserController()
	http.Handle("/hey", c)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

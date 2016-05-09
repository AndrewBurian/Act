package main

import (
	"net/http"
	"log"
)

func main() {

	users := &UsersHandler{}

	http.Handle("/api/users/", users)
	http.Handle("/", http.FileServer(http.Dir("/static")))

	err := http.ListenAndServe(":8282", nil)
	if err != nil {
		log.Fatal(err)
	}
}

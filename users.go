package main

import (
	"io"
	"log"
	"net/http"
)

type UsersHandler struct {
	numConnected int
}

func (h *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println("Entering Users Handler")

	io.WriteString(w, "Sup users")

}

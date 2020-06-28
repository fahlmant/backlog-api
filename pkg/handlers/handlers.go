package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequests(router *mux.Router) {
	router.HandleFunc("/", homePage).Methods("GET")
	registerGameHandlers(router)
}

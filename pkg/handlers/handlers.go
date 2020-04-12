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
	router.HandleFunc("/games", gamesHandler).Methods("GET")
	router.HandleFunc("/games", gamePost).Methods("POST")
	router.HandleFunc("/games/{id}", gameGet).Methods("GET")
	router.HandleFunc("/games/{id}", gamePut).Methods("PUT")
	router.HandleFunc("/games/{id}", gameDelete).Methods("DELETE")
}

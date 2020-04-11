package main

import (
    
    "log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/fahlmant/backlog-api/pkg/handlers"
)

func main() {

	router := mux.NewRouter()
	
	handlers.HandleRequests(router)
	
	http.Handle("/", router)
    log.Fatal(http.ListenAndServe(":8081", nil))
}
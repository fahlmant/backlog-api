package main

import (
	"log"
	"net/http"

	mydb "github.com/fahlmant/backlog-api/pkg/database"
	"github.com/fahlmant/backlog-api/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	handlers.HandleRequests(router)

	mydb.InitDb()
	defer mydb.DB.Close()

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

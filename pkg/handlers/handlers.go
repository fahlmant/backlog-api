package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	types "github.com/fahlmant/backlog-api/pkg/types"
	mydb "github.com/fahlmant/backlog-api/pkg/database"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	fmt.Println("Endpoint Hit: homePage")
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {

	games, err := mydb.GetAllGames(mydb.DB)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func gameGet(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	uuid, err := uuid.Parse(id)
	if err != nil {
		panic(err)
	}
	game, err := mydb.GetGame(mydb.DB, uuid)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}

func gamePost(w http.ResponseWriter, r *http.Request) {

	var game types.Game
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &game)

	if game.Title == "" || game.Platform == "" {
		json.NewEncoder(w).Encode(struct{Error string}{"Missing Title or Platform",})
		return
	}

	mydb.CreateGame(mydb.DB, game)

}

func gamePut(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	uuid, err := uuid.Parse(id)
	if err != nil {
		panic(err)
	}

	var game types.Game
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &game)

	if game.Title == "" || game.Platform == "" {
		json.NewEncoder(w).Encode(struct{Error string}{"Missing Title or Platform",})
		return
	}

	game.ID = uuid

	mydb.UpdateGame(mydb.DB, game)

}

func gameDelete(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	uuid, err := uuid.Parse(id)
	if err != nil {
		panic(err)
	}
	err = mydb.DeleteGame(mydb.DB, uuid)
	if err != nil {
		panic(err)
	}

	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(game)
}

func HandleRequests(router *mux.Router) {
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/games", gamesHandler).Methods("GET")
	router.HandleFunc("/games", gamePost).Methods("POST")
	router.HandleFunc("/games/{id}", gameGet).Methods("GET")
	router.HandleFunc("/games/{id}", gamePut).Methods("PUT")
	router.HandleFunc("/games/{id}", gameDelete).Methods("DELETE")
}

package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	mydb "github.com/fahlmant/backlog-api/pkg/database"
	types "github.com/fahlmant/backlog-api/pkg/types"
)

func registerGameHandlers(router *mux.Router) {
	router.HandleFunc("/games", gamesHandler).Methods("GET")
	router.HandleFunc("/games", gamePost).Methods("POST")
	router.HandleFunc("/games/{id}", gameGet).Methods("GET")
	router.HandleFunc("/games/{id}", gamePut).Methods("PUT")
	router.HandleFunc("/games/{id}", gamePatch).Methods("PATCH")
	router.HandleFunc("/games/{id}", gameDelete).Methods("DELETE")
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {

	//Get all games from database
	//TODO implement pagination
	games, err := mydb.GetAllGames(mydb.DB)
	if err != nil {
		panic(err)
	}

	//Write JSON of all objects
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func gameGet(w http.ResponseWriter, r *http.Request) {

	//Retrieve the object ID from the request and store it in the game struct
	id := mux.Vars(r)["id"]
	uuid, err := uuid.Parse(id)
	if err != nil {
		panic(err)
	}

	//Retrieve game object from database
	game, err := mydb.GetGame(mydb.DB, uuid)

	//Write JSON of object
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}

func gamePost(w http.ResponseWriter, r *http.Request) {

	//Declare types needed
	var game types.Game

	//Build game struct from request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &game)

	//Check that all the required data is in response body
	if game.Title == "" || game.Platform == "" {
		json.NewEncoder(w).Encode(struct{ Error string }{"Missing Title or Platform"})
		return
	}

	//Generate a new uuid
	game.ID = uuid.New()

	//Create entry in database
	err := mydb.CreateGame(mydb.DB, game)
	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{"Error creating new game"})
		return
	}

	//Write JSON of new entry if successful
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
	return
}

func gamePut(w http.ResponseWriter, r *http.Request) {

	//Delcare types needed
	var game types.Game
	var err error

	//Retrieve the object ID from the request and store it in the game struct
	id := mux.Vars(r)["id"]
	game.ID, err = uuid.Parse(id)
	if err != nil {
		panic(err)
	}

	//Get the rest of the game object data from the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &game)

	//For PUT, ensure all fields are filled in
	if game.Title == "" || game.Platform == "" {
		json.NewEncoder(w).Encode(struct{ Error string }{"Missing Title or Platform"})
		return
	}

	//Update the database entry
	mydb.UpdateGame(mydb.DB, game)
}

func gamePatch(w http.ResponseWriter, r *http.Request) {

	//Delcare types needed
	var game types.Game
	//var err error

	//Retrieve the object ID from the request and store it in the game struct
	id := mux.Vars(r)["id"]

	game, _ = mydb.GetGame(mydb.DB, uuid.MustParse(id))

	//Get the rest of the game object data from the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &game)

	if game.Title == "" || game.Platform == "" {
		json.NewEncoder(w).Encode(struct{ Error string }{"Missing Title or Platform"})
		return
	}

	//Update the database entry
	mydb.UpdateGame(mydb.DB, game)

}

func gameDelete(w http.ResponseWriter, r *http.Request) {

	//Retrieve the object ID from the request and store it in the game struct
	id := mux.Vars(r)["id"]
	uuid, err := uuid.Parse(id)
	if err != nil {
		panic(err)
	}

	//Delete database entry
	err = mydb.DeleteGame(mydb.DB, uuid)
	if err != nil {
		panic(err)
	}

}

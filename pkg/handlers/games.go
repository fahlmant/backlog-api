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
		json.NewEncoder(w).Encode(struct{ Error string }{"Missing Title or Platform"})
		return
	}

	uuid := uuid.New()
	game.ID = uuid

	err := mydb.CreateGame(mydb.DB, game)

	if err != nil {
		json.NewEncoder(w).Encode(struct{ Error string }{"Error creating new game"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
	return
}

func gamePut(w http.ResponseWriter, r *http.Request) {

	var game types.Game
	var err error

	id := mux.Vars(r)["id"]
	game.ID, err = uuid.Parse(id)
	if err != nil {
		panic(err)
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &game)

	if game.Title == "" || game.Platform == "" {
		json.NewEncoder(w).Encode(struct{ Error string }{"Missing Title or Platform"})
		return
	}

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
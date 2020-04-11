package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	
	types "github.com/fahlmant/backlog-api/pkg/types"
)


func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
    fmt.Println("Endpoint Hit: homePage")
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {

	var games []types.Game
	game := types.Game{
		ID: uuid.New(),
		Title: "Foo",
		Platform: "Bar",
	}
	game2 := types.Game{
		ID: uuid.New(),
		Title: "Boo",
		Platform: "Far",
	}

	games = append(games, game)
	games = append(games, game2)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func gameGet(w http.ResponseWriter, r *http.Request) {

	//fmt.Printf("%+v\n", mux.Vars(r)["id"])

	game := types.Game{
		ID: uuid.New(),
		Title: "Foo",
		Platform: "Bar",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}

func gamePost(w http.ResponseWriter, r *http.Request) {

	var game types.Game
	//fmt.Fprintf(w, "POST/PUT Single Game: ")
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &game)
	fmt.Printf("%+v\n",game)
}

func gamePut(w http.ResponseWriter, r *http.Request) {

	
}

func gameDelete(w http.ResponseWriter, r *http.Request) {

	
}

func HandleRequests(router *mux.Router) {
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/games", gamesHandler).Methods("GET")
	router.HandleFunc("/games", gamePut).Methods("PUT")
	router.HandleFunc("/games/{id}", gameGet).Methods("GET")
	router.HandleFunc("/games/{id}", gamePost).Methods("POST")
	router.HandleFunc("/games/{id}", gameDelete).Methods("DELETE")
}
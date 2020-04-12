package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/fahlmant/backlog-api/pkg/types"
)

const (
	dbhost    = "DBHOST"
	dbport    = "DBPORT"
	dbuser    = "DBUSER"
	dbpass    = "DBPASS"
	dbname    = "DBNAME"
	gameTable = `games`
)

var (
	DB *sql.DB
)

func InitDb() {
	var err error
	config := dbConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}

func GetAllGames(db *sql.DB) ([]types.Game, error) {

	var games []types.Game
	rows, err := db.Query(`
	SELECT
		id,
		title,
		platform
	FROM games`)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	defer rows.Close()
	for rows.Next() {
		game := types.Game{}
		err = rows.Scan(
			&game.ID,
			&game.Title,
			&game.Platform,
		)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	return games, nil
}

func GetGame(db *sql.DB, uuid uuid.UUID) (types.Game, error) {

	var game types.Game

	sqlStatement := `SELECT * FROM ` + gameTable + ` WHERE id=$1`
	row := db.QueryRow(sqlStatement, uuid)
	err := row.Scan(
		&game.ID,
		&game.Title,
		&game.Platform,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			fmt.Printf("%+v\n", err)
			return types.Game{}, err
		}
	}

	return game, nil
}


func CreateGame(db *sql.DB, game types.Game) error {

	uuid := uuid.New()
	sqlStatement := `INSERT INTO ` + gameTable + ` (ID, Title, Platform)
					VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, uuid, game.Title, game.Platform)
	if err != nil {
		return err
	}

	return nil
}

func UpdateGame(db *sql.DB, game types.Game) error{

	sqlStatement := `UPDATE ` + gameTable + ` SET Title=$2, Platform=$3 WHERE id=$1`
	_, err := db.Exec(sqlStatement, game.ID, game.Title, game.Platform)
	if err != nil {
		return err
	}

	return nil
}

func DeleteGame(db *sql.DB, uuid uuid.UUID) error {


	sqlStatement := `DELETE FROM ` + gameTable + ` WHERE id=$1`
	_, err := db.Exec(sqlStatement, uuid)
	if err != nil {
		return err
	}

	return nil
}
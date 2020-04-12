package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/fahlmant/backlog-api/pkg/types"
)

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

	sqlStatement := `INSERT INTO ` + gameTable + ` (ID, Title, Platform)
					VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, game.ID, game.Title, game.Platform)
	if err != nil {
		return err
	}

	return nil
}

func UpdateGame(db *sql.DB, game types.Game) error {

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

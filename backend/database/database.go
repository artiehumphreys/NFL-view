package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS players (
		"id" INTEGER NOT NULL PRIMARY KEY,
		"play_id" INTEGER,
		"nfl_player_id" INTEGER,
		"type" TEXT,
		"game_position" TEXT,
		"team" TEXT,
		"jersey_number" TEXT,
		"first_name" TEXT,
		"last_name" TEXT,
		"quality" TEXT
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

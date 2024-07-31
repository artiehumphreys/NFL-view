package database

import (
	"database/sql"
	"log"

	"github.com/artiehumphreys/NFL-view/models"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS injuries (
		auto_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"game" INTEGER,
		"play_id" INTEGER,
		"nfl_player_id" INTEGER,
		"type" TEXT,
		"game_position" TEXT,
		"team" TEXT,
		"jersey_number" TEXT,
		"first_name" TEXT,
		"last_name" TEXT,
		"quality" TEXT,
		UNIQUE (game, play_id, nfl_player_id)
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func PopulateDB(db *sql.DB, records []models.Record) {
	var i = 0
	for {
		if i == len(records) {
			break
		}
		record := records[i]
		preparedInsertRecord := `INSERT OR IGNORE INTO injuries (game, play_id, nfl_player_id, type, game_position, team, jersey_number, first_name, last_name, quality) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		statement, err := db.Prepare(preparedInsertRecord)
		if err != nil {
			log.Fatal(err)
		}
		_, err = statement.Exec(record.ID, record.PlayID, record.NFLPlayerID, record.Type, record.GamePosition, record.Team, record.JerseyNumber, record.FirstName, record.LastName, record.Quality)
		if err != nil {
			log.Fatal(err)
		}
		i += 1
	}
}

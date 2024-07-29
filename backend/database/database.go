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

	createTableSQL := `CREATE TABLE IF NOT EXISTS players (
		auto_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"id" INTEGER,
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

func PopulateDB(db *sql.DB, records []models.Record) {
	var i = 0
	for {
		if i == len(records) {
			break
		}
		record := records[i]
		preparedInsertRecord := `INSERT INTO players(id, play_id, nfl_player_id, type, game_position, team, jersey_number, first_name, last_name, quality) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
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

func GetSearchTags(db *sql.DB) ([]string, error) {
	var query = "SELECT type FROM players"
	rows, err := db.Query(query)
	var set = models.NewSet()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, err
		}
		set.Add(tag)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return set.ToSlice(), nil
}

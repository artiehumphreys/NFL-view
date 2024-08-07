package database

import (
	"database/sql"
	"fmt"
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
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"game" TEXT,
		"play_id" TEXT,
		"nfl_player_id" TEXT,
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

func IsTablePopulated(db *sql.DB) bool {
	var count int
	query := "SELECT COUNT(*) FROM injuries"
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}

func PopulateDB(db *sql.DB, records []models.Record) {
	var i = 0
	for {
		if i == len(records) {
			break
		}
		record := &records[i]
		findDuplicates(db, record)
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

func findDuplicates(db *sql.DB, record *models.Record) {
	var existingCount int
	query := "SELECT COUNT(*) FROM injuries WHERE game = ? AND play_id = ?"
	err := db.QueryRow(query, record.ID, record.PlayID).Scan(&existingCount)
	if err != nil {
		log.Fatal(err)
	}

	count := 1
	originalPlayID := record.PlayID
	fmt.Println(originalPlayID)

	for existingCount > 0 {
		newPlayID := fmt.Sprintf("%s_%d", originalPlayID, count)
		if count == 1 {
			updateQuery := "UPDATE injuries SET play_id = ? WHERE game = ? AND play_id = ?"
			_, err = db.Exec(updateQuery, newPlayID, record.ID, record.PlayID)
			if err != nil {
				log.Fatal(err)
			}
		}

		record.PlayID = newPlayID
		count++
		err = db.QueryRow(query, record.ID, record.PlayID).Scan(&existingCount)
		if err != nil {
			log.Fatal(err)
		}
	}
}

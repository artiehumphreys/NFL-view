package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/artiehumphreys/NFL-view/models"
)

func GetField(db *sql.DB, field string) ([]string, error) {
	var query = fmt.Sprintf("SELECT %s FROM injuries", field)
	rows, err := db.Query(query)
	var set = models.NewSet()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var instance string
		if err := rows.Scan(&instance); err != nil {
			return nil, err
		}
		set.Add(instance)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return set.ToSlice(), nil
}

func GetInjuryInfo(db *sql.DB) ([]models.InjuryDisplay, error) {
	var query = "SELECT game, play_id, type, game_position, team, jersey_number, first_name, last_name FROM injuries"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.InjuryDisplay
	for rows.Next() {
		var instance models.InjuryDisplay
		if err := rows.Scan(&instance.Game, &instance.PlayID, &instance.Type, &instance.GamePosition, &instance.Team, &instance.JerseyNumber, &instance.FirstName, &instance.LastName); err != nil {
			return nil, err
		}
		results = append(results, instance)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func GetGamesList(db *sql.DB) (map[models.GamePlayKey][]string, error) {
	var query = "SELECT game, play_id, team, first_name, last_name FROM injuries"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	results := make(map[models.GamePlayKey][]string)
	for rows.Next() {
		var instance models.GameDisplay
		if err := rows.Scan(&instance.Game, &instance.PlayID, &instance.Team, &instance.FirstName, &instance.LastName); err != nil {
			return nil, err
		}
		key := models.GamePlayKey{
			Game:   instance.Game,
			PlayID: instance.PlayID,
		}
		results[key] = append(results[key], instance.String())
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func RemoveInjury(db *sql.DB, injury models.InjuryDisplay) error {
	query := `
		DELETE FROM injuries 
		WHERE game = ? 
		AND play_id = ?
		AND first_name = ? 
		AND last_name = ? 
		AND type = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(injury.Game, injury.PlayID, injury.FirstName, injury.LastName, injury.Type)
	if err != nil {
		return err
	}
	return nil
}
func GetGameInfo(db *sql.DB, gameID string) ([]models.InjuryDisplay, error) {
	count := 1

	replace := func(id *models.InjuryDisplay) {
		id.PlayID = id.PlayID + "_" + strconv.Itoa(count)
	}

	findAndReplace := func(list []models.InjuryDisplay, toFind string) bool {
		flag := false
		for i := range list {
			if list[i].PlayID == toFind {
				replace(&list[i])
				count++
				flag = true
			}
		}
		return flag
	}

	query := "SELECT game, play_id, type, game_position, team, jersey_number, first_name, last_name FROM injuries WHERE game = ?"
	rows, err := db.Query(query, gameID)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.InjuryDisplay
	for rows.Next() {
		var instance models.InjuryDisplay
		if err := rows.Scan(&instance.Game, &instance.PlayID, &instance.Type, &instance.GamePosition, &instance.Team, &instance.JerseyNumber, &instance.FirstName, &instance.LastName); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		if findAndReplace(results, instance.PlayID) {
			replace(&instance)
		}
		results = append(results, instance)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Row iteration error: %v", err)
		return nil, err
	}

	return results, nil
}

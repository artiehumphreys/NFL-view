package database

import (
	"database/sql"
	"fmt"

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

func GetInjuryInfo(db *sql.DB) ([]string, error) {
	var query = "SELECT game, type, game_position, team, jersey_number, first_name, last_name FROM injuries"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []string
	for rows.Next() {
		var instance models.InjuryDisplay
		if err := rows.Scan(&instance.Game, &instance.Type, &instance.GamePosition, &instance.Team, &instance.JerseyNumber, &instance.FirstName, &instance.LastName); err != nil {
			return nil, err
		}
		results = append(results, instance.String())
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func GetGameInfo(db *sql.DB) (map[string][]string, error) {
	var query = "SELECT game, play_id, team, first_name, last_name FROM injuries"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	results := make(map[string][]string)
	for rows.Next() {
		var instance models.GameDisplay
		if err := rows.Scan(&instance.Game, &instance.PlayID, &instance.Team, &instance.FirstName, &instance.LastName); err != nil {
			return nil, err
		}
		results[instance.Game] = append(results[instance.Game], instance.String())
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func RemoveInjury(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM players WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

package database

import (
	"database/sql"

	"github.com/artiehumphreys/NFL-view/models"
)

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

func GetInjuryInfo(db *sql.DB) ([]string, error) {
	var query = "SELECT id, type, game_position, team, jersey_number, first_name, last_name FROM players"
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

func GetGameInfo(db *sql.DB) ([]string, error) {
	var query = "SELECT id, first_name, last_name FROM players"
	rows, err := db.Query(query)
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
}

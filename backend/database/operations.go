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

func GetDisplayInfo(db *sql.DB) ([]models.InjuryDisplay, error) {
	var query = "SELECT type, game_position, team, jersey_number, first_name, last_name FROM players"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var info []models.InjuryDisplay
	for rows.Next() {
		var instance models.InjuryDisplay
		if err := rows.Scan(&instance.Type, &instance.GamePosition, &instance.Team, &instance.JerseyNumber, &instance.FirstName, &instance.LastName); err != nil {
			return nil, err
		}
		info = append(info, instance)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return info, nil

}

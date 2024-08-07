package database

import (
	"database/sql"

	"github.com/artiehumphreys/NFL-view/models"
)

func SearchDB(db *sql.DB, field string, keyword string) ([]models.InjuryDisplay, error) {
	query := "SELECT game, play_id, type, game_position, team, jersey_number, first_name, last_name FROM injuries WHERE " + field + " LIKE ?"
	rows, err := db.Query(query, "%"+keyword+"%")
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

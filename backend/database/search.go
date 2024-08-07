package database

import (
	"database/sql"
	"sync"

	"github.com/artiehumphreys/NFL-view/models"
)

func searchDB(db *sql.DB, field string, keyword string) ([]models.InjuryDisplay, error) {
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

func ConcurrentSearch(db *sql.DB, keyword string) ([]models.InjuryDisplay, error) {
	var wg sync.WaitGroup
	resultsChan := make(chan []models.InjuryDisplay, 4)
	errChan := make(chan error, 4)

	fields := []string{"first_name", "last_name", "team", "type"}

	for _, field := range fields {
		wg.Add(1)
		go func(field string) {
			defer wg.Done()
			results, err := searchDB(db, field, keyword)
			if err != nil {
				errChan <- err
				return
			}
			resultsChan <- results
		}(field)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
		close(errChan)
	}()

	var finalResults []models.InjuryDisplay
	for results := range resultsChan {
		finalResults = append(finalResults, results...)
	}

	if err := <-errChan; err != nil {
		return nil, err
	}

	return finalResults, nil
}

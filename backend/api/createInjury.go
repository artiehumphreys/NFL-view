package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/artiehumphreys/NFL-view/models"
	"github.com/julienschmidt/httprouter"
)

func CreateInjuryHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var injury models.Record
		if err := json.NewDecoder(r.Body).Decode(&injury); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		query := `
			INSERT INTO injuries (play_id, nfl_player_id, type, game_position, team, jersey_number, first_name, last_name, quality)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
		stmt, err := db.Prepare(query)
		if err != nil {
			http.Error(w, "Failed to prepare insert query", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		result, err := stmt.Exec(
			injury.PlayID,
			injury.NFLPlayerID,
			injury.Type,
			injury.GamePosition,
			injury.Team,
			injury.JerseyNumber,
			injury.FirstName,
			injury.LastName,
			injury.Quality,
		)
		if err != nil {
			http.Error(w, "Failed to execute insert query", http.StatusInternalServerError)
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, "Failed to retrieve last insert id", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"id": %d}`, id)))
	}
}

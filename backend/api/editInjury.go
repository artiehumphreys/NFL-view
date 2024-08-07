package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/artiehumphreys/NFL-view/models"
	"github.com/julienschmidt/httprouter"
)

func EditInjuryHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")
		if id == "" {
			http.Error(w, "Missing ID in request URL", http.StatusBadRequest)
			return
		}

		var injury models.Record
		if err := json.NewDecoder(r.Body).Decode(&injury); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		fmt.Println(&injury)

		query := `
			UPDATE injuries 
			SET play_id = ?, nfl_player_id = ?, type = ?, game_position = ?, team = ?, jersey_number = ?, first_name = ?, last_name = ?, quality = ?
			WHERE id = ?`
		stmt, err := db.Prepare(query)
		if err != nil {
			http.Error(w, "Failed to prepare update query", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(
			injury.PlayID,
			injury.NFLPlayerID,
			injury.Type,
			injury.GamePosition,
			injury.Team,
			injury.JerseyNumber,
			injury.FirstName,
			injury.LastName,
			injury.Quality,
			id,
		)
		if err != nil {
			http.Error(w, "Failed to execute update query", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

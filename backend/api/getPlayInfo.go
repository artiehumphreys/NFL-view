package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/artiehumphreys/NFL-view/database"
	"github.com/julienschmidt/httprouter"
)

func GetPlayInfoHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		gameID := ps.ByName("gameId")
		playID := ps.ByName("playId")

		play, err := database.GetPlayInfo(db, gameID, playID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(play)
	}
}

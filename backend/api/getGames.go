package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sort"

	"github.com/artiehumphreys/NFL-view/database"
	"github.com/julienschmidt/httprouter"
)

func GetGamesHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		games, err := database.GetField(db, "game")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sort.Strings(games)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(games)
	}
}
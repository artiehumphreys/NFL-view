package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sort"

	"github.com/artiehumphreys/NFL-view/database"
	"github.com/julienschmidt/httprouter"
)

func GetGamesListHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		gameInfo, err := database.GetGamesList(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		keys := make([]string, 0, len(gameInfo))
		for k := range gameInfo {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		sorted := make([]GameInfo, 0, len(keys))

		for _, k := range keys {
			sorted = append(sorted, GameInfo{
				Game:   k,
				Events: gameInfo[k],
			})
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(sorted)
	}
}

type GameInfo struct {
	Game   string   `json:"game"`
	Events []string `json:"events"`
}

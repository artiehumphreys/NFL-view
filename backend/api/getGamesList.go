package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sort"

	"github.com/artiehumphreys/NFL-view/database"
	"github.com/artiehumphreys/NFL-view/models"
	"github.com/julienschmidt/httprouter"
)

func GetGamesListHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		gameList, err := database.GetGamesList(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		keys := make([]models.GamePlayKey, 0, len(gameList))
		for k := range gameList {
			keys = append(keys, k)
		}

		sort.Slice(keys, func(i, j int) bool {
			if keys[i].Game == keys[j].Game {
				return keys[i].PlayID < keys[j].PlayID
			}
			return keys[i].Game < keys[j].Game
		})

		sorted := make([]GameList, 0, len(keys))

		for _, k := range keys {
			sorted = append(sorted, GameList{
				Game:   k.Game,
				PlayID: k.PlayID,
				Events: gameList[k],
			})
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(sorted)
	}
}

type GameList struct {
	Game   string   `json:"game"`
	PlayID string   `json:"play_id"`
	Events []string `json:"events"`
}

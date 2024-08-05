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

		playList := make(map[string][]string)
		for _, k := range keys {
			playList[k.Game] = append(playList[k.Game], k.PlayID)
		}

		res := make([]GameList, 0, len(playList))
		for game, playIDs := range playList {
			var events []string
			for _, k := range keys {
				if k.Game == game {
					events = append(events, gameList[k]...)
				}
			}
			res = append(res, GameList{
				Game:   game,
				PlayID: playIDs,
				Events: events,
			})
		}

		sort.Slice(res, func(i, j int) bool {
			return res[i].Game < res[j].Game
		})

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(res)
	}
}

type GameList struct {
	Game   string   `json:"game"`
	PlayID []string `json:"play_id"`
	Events []string `json:"events"`
}

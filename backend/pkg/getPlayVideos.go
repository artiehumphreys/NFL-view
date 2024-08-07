package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func GetPlayVideosHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		gameID := ps.ByName("gameId")
		playID := ps.ByName("playId")

		split := strings.Split(playID, "_")
		if strings.Contains(playID, "_") {
			playID = split[0]
		}

		formattedPlayID := fmt.Sprintf("%06s", playID)
		formattedGameID := fmt.Sprintf("%06s", gameID)

		patterns := []string{
			fmt.Sprintf("%s_%s_0002_000_*.mp4", formattedGameID, formattedPlayID),
			fmt.Sprintf("%s_%s_0006_000_*.mp4", formattedGameID, formattedPlayID),
		}

		var videos []string
		for _, glob := range patterns {
			pattern := filepath.Join("../public/alpha/nfl_videos", glob)
			files, err := filepath.Glob(pattern)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			for _, file := range files {
				filename := filepath.Base(file)
				videos = append(videos, filename)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(videos)
	}
}

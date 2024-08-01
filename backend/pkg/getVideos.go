package pkg

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func GetVideosHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		query := r.URL.Query()
		gameID := query.Get("game")

		pattern := filepath.Join("../../alpha/nfl_videos", "0"+gameID+"*0001_000_15000.mp4")
		files, err := filepath.Glob(pattern)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var videos []string
		for _, file := range files {
			videos = append(videos, filepath.Base(file))
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(videos)
	}
}

package pkg

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func GetVideosHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		gameID := ps.ByName("id")

		patterns := []string{
			"0" + gameID + "_*_0001_000_*.mp4",
			"0" + gameID + "_*_0002_000_*.mp4",
			"0" + gameID + "_*_0006_000_*.mp4",
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

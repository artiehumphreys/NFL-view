package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sort"

	"github.com/artiehumphreys/NFL-view/database"
	"github.com/julienschmidt/httprouter"
)

func GetSearchTagsHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		searchTags, err := database.GetSearchTags(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sort.Strings(searchTags)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(searchTags)
	}
}

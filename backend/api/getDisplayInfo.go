package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/artiehumphreys/NFL-view/database"
	"github.com/julienschmidt/httprouter"
)

func GetDisplayInfoHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		displayInfo, err := database.GetDisplayInfo(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		slices.Sort(displayInfo)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(displayInfo)
	}
}

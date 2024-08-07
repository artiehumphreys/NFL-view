package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/artiehumphreys/NFL-view/database"
	"github.com/artiehumphreys/NFL-view/models"
	"github.com/julienschmidt/httprouter"
)

func GetDisplayInfoHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		queryParams := r.URL.Query()
		search := queryParams.Get("search")
		tag := queryParams.Get("tag")
		var injuries []models.InjuryDisplay
		var err error

		switch {
		case search != "":
			//TODO: Search logic
			injuries = make([]models.InjuryDisplay, 0, 1)
		case tag != "":
			injuries, err = database.GetInjuriesByType(db, tag)
		default:
			injuries, err = database.GetInjuryInfo(db)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(injuries)
	}
}

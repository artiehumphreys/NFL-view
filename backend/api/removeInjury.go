package api

import (
	"database/sql"
	"net/http"

	"github.com/artiehumphreys/NFL-view/database"
	"github.com/artiehumphreys/NFL-view/models"
	"github.com/julienschmidt/httprouter"
)

func RemoveInjuryHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		query := r.URL.Query()
		game := query.Get("game")
		play_id := query.Get("play_id")
		fName := query.Get("fName")
		lName := query.Get("lName")
		typeStr := query.Get("type")
		injury := models.InjuryDisplay{
			Game:         game,
			PlayID:       play_id,
			GamePosition: "",
			Team:         "",
			JerseyNumber: "",
			FirstName:    fName,
			LastName:     lName,
			Type:         typeStr,
		}

		err := database.RemoveInjury(db, injury)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Item successfully deleted"))
	}
}

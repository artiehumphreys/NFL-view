package api

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/artiehumphreys/NFL-view/database"
	"github.com/julienschmidt/httprouter"
)

func RemoveInjuryHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")
		items := strings.Split(id, "%20")

		err := database.RemoveInjury(db, items)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Item successfully deleted"))
	}
}

package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetGameInfoHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Println("hi")
	}
}

package route

import (
	"database/sql"

	"github.com/artiehumphreys/NFL-view/api"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, db *sql.DB) {
	router.GET("/tags", api.GetSearchTagsHandler(db))
	router.GET("/info", api.GetDisplayInfoHandler(db))
	router.GET("/games", api.GetGameInfoHandler(db))
}

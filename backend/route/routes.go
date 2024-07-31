package route

import (
	"database/sql"

	"github.com/artiehumphreys/NFL-view/api"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, db *sql.DB) {
	router.GET("/home/tags", api.GetSearchTagsHandler(db))
	router.GET("/home/info", api.GetDisplayInfoHandler(db))
	router.GET("/home/games", api.GetGameInfoHandler(db))
}

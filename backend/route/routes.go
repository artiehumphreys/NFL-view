package route

import (
	"database/sql"

	"github.com/artiehumphreys/NFL-view/api"
	"github.com/artiehumphreys/NFL-view/pkg"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, db *sql.DB) {
	router.GET("/tags", api.GetSearchTagsHandler(db))
	router.GET("/info", api.GetDisplayInfoHandler(db))
	router.GET("/games", api.GetGamesListHandler(db))
	router.GET("/game", api.GetGameInfoHandler(db))
	router.GET("/video", pkg.GetVideosHandler())
	router.DELETE("/removeInjury", api.RemoveInjuryHandler(db))
}

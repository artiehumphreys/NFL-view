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
	router.GET("/games/:id", api.GetGameInfoHandler(db))
	router.GET("/games/:id/videos", pkg.GetGameVideosHandler())
	router.GET("/plays/:id/videos", pkg.GetPlayVideosHandler())
	router.DELETE("/injuries", api.RemoveInjuryHandler(db))
}

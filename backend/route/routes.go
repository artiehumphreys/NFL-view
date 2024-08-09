package route

import (
	"database/sql"

	"github.com/artiehumphreys/NFL-view/api"
	"github.com/artiehumphreys/NFL-view/pkg"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, db *sql.DB) {
	router.GET("/tags", api.GetSearchTagsHandler(db))
	router.GET("/injuries", api.GetDisplayInfoHandler(db))
	router.GET("/games", api.GetGamesListHandler(db))
	router.GET("/games/:gameId", api.GetGameInfoHandler(db))
	router.GET("/games/:gameId/videos", pkg.GetGameVideosHandler())
	router.GET("/games/:gameId/plays/:playId", api.GetPlayInfoHandler(db))
	router.GET("/games/:gameId/plays/:playId/videos", pkg.GetPlayVideosHandler())
	router.POST("/games/:gameId/plays/:playId/conversions", pkg.ConvertPlayHandler())
	router.DELETE("/injuries", api.RemoveInjuryHandler(db))
	router.PUT("/injuries/:id", api.EditInjuryHandler(db))
	router.POST("/injuries", api.CreateInjuryHandler(db))
}

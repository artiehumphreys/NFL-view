package main

import (
	"net/http"

	"github.com/artiehumphreys/NFL-view/api"
	"github.com/artiehumphreys/NFL-view/database"
	"github.com/artiehumphreys/NFL-view/pkg"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	db := database.InitDB("../NFLView.db")
	defer db.Close()

	filePath := "../alpha/NFLview_Database_2024.07.csv"
	records := pkg.ReadCsvFile(filePath)

	database.PopulateDB(db, records)

	router := httprouter.New()
	router.GET("/home", api.GetSearchTagsHandler(records))

	corsHandler := cors.Default().Handler(router)

	http.ListenAndServe(":8080", corsHandler)
}

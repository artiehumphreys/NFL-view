package main

import (
	"fmt"
	"net/http"

	"github.com/artiehumphreys/NFL-view/database"
	"github.com/artiehumphreys/NFL-view/pkg"
	"github.com/artiehumphreys/NFL-view/route"
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

	route.RegisterRoutes(router, db)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	fmt.Println("Listening...")
	_ = http.ListenAndServe(":8080", corsHandler)
}

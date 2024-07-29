package main

import (
	"fmt"

	"github.com/artiehumphreys/NFL-view/database"
	"github.com/artiehumphreys/NFL-view/pkg"
)

func main() {
	db := database.InitDB("../NFLView.db")
	defer db.Close()

	filePath := "../alpha/NFLview_Database_2024.07.csv"
	records := pkg.ReadCsvFile(filePath)

	database.PopulateDB(db, records)

	for _, record := range records {
		fmt.Printf("%+v\n", record)
	}
}

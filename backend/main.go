package main

import (
	"fmt"

	"github.com/artiehumphreys/NFL-view/pkg"
)

func main() {
	filePath := "../alpha/NFLview_Database_2024.07.csv"
	records := pkg.ReadCsvFile(filePath)

	for _, record := range records {
		fmt.Printf("%+v\n", record)
	}
}

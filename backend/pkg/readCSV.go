package pkg

import (
	"encoding/csv"
	"log"
	"os"
)

func readCsvFile(filePath string) {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)

    
}
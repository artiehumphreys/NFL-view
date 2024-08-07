package pkg

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/artiehumphreys/NFL-view/models"
)

func ReadCsvFile(filePath string) []models.Record {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	return csvToRecord(*csvReader)
}

func csvToRecord(reader csv.Reader) []models.Record {
	var records []models.Record
	var header []string
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(header) == 0 {
			header = row
		} else {
			records = append(records, models.Record{
				ID:           row[0],
				PlayID:       row[1],
				NFLPlayerID:  row[2],
				Type:         row[3],
				GamePosition: row[4],
				Team:         row[5],
				JerseyNumber: row[6],
				FirstName:    row[7],
				LastName:     row[8],
				Quality:      row[9],
			})
		}
	}
	return records
}

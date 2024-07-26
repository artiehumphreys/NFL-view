package pkg

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

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
			id, err := strconv.Atoi(row[0])
			if err != nil {
				log.Fatal(err)
			}
			playID, err := strconv.Atoi(row[1])
			if err != nil {
				log.Fatal(err)
			}
			nflPlayID, err := strconv.Atoi(row[2])
			if err != nil {
				log.Fatal(err)
			}
			records = append(records, models.Record{
				ID:           id,
				PlayID:       playID,
				NFLPlayerID:  nflPlayID,
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

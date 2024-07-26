package pkg

import (
	"encoding/csv"
	"io"
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
    data := csvToMap(*csvReader)
    
}

func csvToMap(reader csv.Reader) []map[string]string {
    rows := []map[string]string{}
	var header []string
    for {
        row, err := reader.Read()
        if err == io.EOF{
            break
        }
        if err != nil {
			log.Fatal(err)
		}
        if len(header) == 0 {
			header = row
        } else {
            dict := map[string]string{}
            for i := range header {
				dict[header[i]] = row[i]
			}
            rows = append(rows, dict)
        }
    }
    return rows
}
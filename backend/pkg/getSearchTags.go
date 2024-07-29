package pkg

import "github.com/artiehumphreys/NFL-view/models"

func getSearchTags(records []models.Record) []string {
	set := models.NewSet()
	for _, record := range records {
		set.Add(record.Type)
	}
	return set.ToSlice()
}

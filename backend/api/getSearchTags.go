package api

import (
	"encoding/json"
	"net/http"

	"github.com/artiehumphreys/NFL-view/models"
	"github.com/julienschmidt/httprouter"
)

func getSearchTags(records []models.Record) []string {
	set := models.NewSet()
	for _, record := range records {
		set.Add(record.Type)
	}
	return set.ToSlice()
}

func GetSearchTagsHandler(records []models.Record) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		searchTags := getSearchTags(records)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(searchTags)
	}
}

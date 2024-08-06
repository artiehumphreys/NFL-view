package pkg

import (
	"strconv"

	"github.com/artiehumphreys/NFL-view/models"
)

func FindAndReplace(list []models.InjuryDisplay, instance *models.InjuryDisplay, count int) int {

	replace := func(id *models.InjuryDisplay) {
		id.PlayID = id.PlayID + "_" + strconv.Itoa(count)
	}

	for i := range list {
		if list[i].PlayID == instance.PlayID {
			if count == 1 {
				replace(instance)
			}
			count++
			replace(&list[i])
		}
	}
	return count
}

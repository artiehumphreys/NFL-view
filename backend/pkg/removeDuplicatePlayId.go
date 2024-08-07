package pkg

import (
	"fmt"
	"strconv"

	"github.com/artiehumphreys/NFL-view/models"
)

func FindAndReplace(list []models.InjuryDisplay, instance *models.InjuryDisplay, count int) int {

	replace := func(id *models.InjuryDisplay) {
		id.PlayID = id.PlayID + "_" + strconv.Itoa(count)
	}

	for i := range list {
		if list[i].PlayID == instance.PlayID && list[i] != *instance {
			fmt.Println(count, list[i], *instance)
			replace(&list[i])
			count++
			if count == 2 {
				replace(instance)
			}
		}
	}
	return count
}

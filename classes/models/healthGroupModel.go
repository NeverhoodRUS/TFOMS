package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"

	"github.com/mitchellh/mapstructure"
)

func HealthGroupDictionary() ([]entities.HealthGroup, string) {
	maps, err := databaseworkers.GetAllRowsAsMap(healthGroupableName)
	if err != nil {
		return nil, err.Error()
	}
	var result []entities.HealthGroup
	for _, el := range *maps {
		var ent entities.HealthGroup
		mapstructure.Decode(el, &ent)
		result = append(result, ent)
	}
	return result, ""
}

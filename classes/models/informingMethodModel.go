package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"

	"github.com/mitchellh/mapstructure"
)

func InformingMethodDictionary() ([]entities.InformingMethod, string) {
	maps, err := databaseworkers.GetAllRowsAsMap(informingMethodTableName)
	if err != nil {
		return nil, err.Error()
	}
	var result []entities.InformingMethod
	for _, el := range *maps {
		var ent entities.InformingMethod
		mapstructure.Decode(el, &ent)
		result = append(result, ent)
	}
	return result, ""
}

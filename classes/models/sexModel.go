package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"

	"github.com/mitchellh/mapstructure"
)

func SexDictionary() ([]entities.Sex, string) {
	maps, err := databaseworkers.GetAllRowsAsMap(sexTableName)
	if err != nil {
		return nil, err.Error()
	}
	var result []entities.Sex
	for _, el := range *maps {
		var ent entities.Sex
		mapstructure.Decode(el, &ent)
		result = append(result, ent)
	}
	return result, ""
}

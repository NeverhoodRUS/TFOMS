package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"

	"github.com/mitchellh/mapstructure"
)

func InformingTypeDictionary() ([]entities.InformingType, string) {
	maps, err := databaseworkers.GetAllRowsAsMap(informingTypeTableName)
	if err != nil {
		return nil, err.Error()
	}
	var result []entities.InformingType
	for _, el := range *maps {
		var ent entities.InformingType
		mapstructure.Decode(el, &ent)
		result = append(result, ent)
	}
	return result, ""
}

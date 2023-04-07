package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"

	"github.com/mitchellh/mapstructure"
)

func EventTypeDictionary() ([]entities.EventType, string) {
	maps, err := databaseworkers.GetAllRowsAsMap(eventTypeTableName)
	if err != nil {
		return nil, err.Error()
	}
	var result []entities.EventType
	for _, el := range *maps {
		var ent entities.EventType
		mapstructure.Decode(el, &ent)
		result = append(result, ent)
	}
	return result, ""
}

package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"

	"github.com/mitchellh/mapstructure"
)

func ProrityGroupDictionary() ([]entities.PriorityGroup, string) {
	maps, err := databaseworkers.GetAllRowsAsMap(priorityGroupTableName)
	if err != nil {
		return nil, err.Error()
	}
	var result []entities.PriorityGroup
	for _, el := range *maps {
		var ent entities.PriorityGroup
		mapstructure.Decode(el, &ent)
		result = append(result, ent)
	}
	return result, ""
}

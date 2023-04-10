package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"

	"github.com/blockloop/scan"
)

func ProrityGroupDictionary() ([]entities.PriorityGroup, string) {
	rows, err := databaseworkers.GetAllRowsAsMap(priorityGroupTableName)
	if err != nil {
		return nil, err.Error()
	}
	var result []entities.PriorityGroup
	err = scan.Rows(&result, rows)
	if err != nil {
		return nil, err.Error()
	}
	return result, ""
}

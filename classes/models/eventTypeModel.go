package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"
	"tfoms_server/static/names"

	"github.com/blockloop/scan"
)

func EventTypeDictionary() ([]entities.EventType, string) {
	rows, err := databaseworkers.GetAllRows(names.EventTypeTableName)
	if err != nil {
		return nil, err.Error()
	}
	var result []entities.EventType
	err = scan.Rows(&result, rows)
	if err != nil {
		return nil, err.Error()
	}
	return result, ""
}

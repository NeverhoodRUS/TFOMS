package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"

	"github.com/blockloop/scan"
)

func InformingMethodDictionary() ([]entities.InformingMethod, string) {
	rows, err := databaseworkers.GetAllRowsAsMap(informingMethodTableName)
	if err != nil {
		return nil, err.Error()
	}
	var result []entities.InformingMethod
	err = scan.Rows(&result, rows)
	if err != nil {
		return nil, err.Error()
	}
	return result, ""
}

package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"

	"github.com/blockloop/scan"
)

func InformingTypeDictionary() ([]entities.InformingType, string) {
	rows, err := databaseworkers.GetAllRowsAsMap(informingTypeTableName)
	if err != nil {
		return nil, err.Error()
	}
	var result []entities.InformingType
	err = scan.Rows(&result, rows)
	if err != nil {
		return nil, err.Error()
	}
	return result, ""
}

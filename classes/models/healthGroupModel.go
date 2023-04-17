package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"
	"tfoms_server/static/names"

	"github.com/blockloop/scan"
)

func HealthGroupDictionary() ([]entities.HealthGroup, string) {
	rows, err := databaseworkers.GetAllRows(names.HealthGroupTableName)
	if err != nil {
		return nil, err.Error()
	}
	var result []entities.HealthGroup
	err = scan.Rows(&result, rows)
	if err != nil {
		return nil, err.Error()
	}
	return result, ""
}

package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"
	"tfoms_server/static/strings"

	"github.com/blockloop/scan"
)

func SexDictionary() ([]entities.Sex, string) {
	rows, err := databaseworkers.GetAllRows(strings.SexTableName)
	if err != nil {
		return nil, err.Error()
	}
	var result []entities.Sex
	err = scan.Rows(&result, rows)
	if err != nil {
		return nil, err.Error()
	}
	return result, ""
}

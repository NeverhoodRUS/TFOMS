package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"
	"tfoms_server/static/names"

	"github.com/blockloop/scan"
)

func OrganizationDictionary() ([]entities.Organization, string) {
	rows, err := databaseworkers.GetAllRows(names.MedicalOrgTableName)
	if err != nil {
		return nil, err.Error()
	}
	var result []entities.Organization
	err = scan.Rows(&result, rows)
	if err != nil {
		return nil, err.Error()
	}
	return result, ""
}

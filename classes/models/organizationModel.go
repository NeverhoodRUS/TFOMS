package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"

	"github.com/blockloop/scan"
)

func OrganizationDictionary() ([]entities.Organization, string) {
	rows, err := databaseworkers.GetAllRowsAsMap(medicalOrgTableName)
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

package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"

	"github.com/blockloop/scan"
)

func InsuranceOrganizationDictionary() ([]entities.InsuranceOrganization, string) {
	rows, err := databaseworkers.GetAllRowsAsMap(insuranceOrgTableName)
	if err != nil {
		return nil, err.Error()
	}
	var result []entities.InsuranceOrganization
	err = scan.Rows(&result, rows)
	if err != nil {
		return nil, err.Error()
	}
	return result, ""
}

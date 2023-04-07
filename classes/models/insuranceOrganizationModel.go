package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"

	"github.com/mitchellh/mapstructure"
)

func InsuranceOrganizationDictionary() ([]entities.InsuranceOrganization, string) {
	maps, err := databaseworkers.GetAllRowsAsMap(insuranceOrgTableName)
	if err != nil {
		return nil, err.Error()
	}
	var result []entities.InsuranceOrganization
	for _, el := range *maps {
		var ent entities.InsuranceOrganization
		mapstructure.Decode(el, &ent)
		result = append(result, ent)
	}
	return result, ""
}

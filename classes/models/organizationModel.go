package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"

	"github.com/mitchellh/mapstructure"
)

func OrganizationDictionary() ([]entities.Organization, string) {
	maps, err := databaseworkers.GetAllRowsAsMap(medicalOrgTableName)
	if err != nil {
		return nil, err.Error()
	}
	var result []entities.Organization
	for _, el := range *maps {
		var ent entities.Organization
		mapstructure.Decode(el, &ent)
		result = append(result, ent)
	}
	return result, ""
}

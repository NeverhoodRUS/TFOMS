package services

import (
	"tfoms_server/classes/entities"
	"tfoms_server/classes/models"
)

func OrganizationDictionary() ([]entities.Organization, string) {
	return models.OrganizationDictionary()
}

func EventTypeDictionary() ([]entities.EventType, string) {
	return models.EventTypeDictionary()
}

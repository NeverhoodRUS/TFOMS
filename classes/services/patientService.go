package services

import (
	"tfoms_server/classes/entities"
	"tfoms_server/classes/models"
)

func GetPatientJournal(filters map[string]string) ([]entities.Patient, string) {
	return models.GetPatientJournal(filters)
}

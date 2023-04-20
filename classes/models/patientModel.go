package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"
)

func GetPatientJournal(filters map[string]string, limit int) ([]*entities.Patient, string) {
	rows, err := databaseworkers.GetPatientJournal(filters, limit)
	if err != nil {
		return nil, err.Error()
	}
	defer rows.Close()
	var result []*entities.Patient
	for rows.Next() {
		var patientId int
		rows.Scan(&patientId)
		patient := entities.NewPatient(patientId, false)
		result = append(result, patient)
	}
	return result, ""
}

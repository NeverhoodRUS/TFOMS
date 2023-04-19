package models

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/classes/entities"
)

func GetPatientJournal(filters map[string]string) ([]entities.Patient, string) {
	rows, err := databaseworkers.GetPatientJournal(filters)
	if err != nil {
		return nil, err.Error()
	}
	defer rows.Close()
	var result []entities.Patient
	for rows.Next() {
		var id int
		rows.Scan(&id)
		patient := entities.PatientInit(id, false)
		result = append(result, *patient)
	}
	return result, ""
}

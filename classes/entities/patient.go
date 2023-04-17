package entities

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
)

type Patient struct {
	PatientDB     *databaseworkers.PatientColumns
	PlannedYearDB *databaseworkers.PlanningYearColumns
}

func PatientInit(patientId int) *Patient {
	patient := &Patient{PatientDB: &databaseworkers.PatientColumns{}, PlannedYearDB: &databaseworkers.PlanningYearColumns{}}
	patient.PatientDB.GetRowStruct(patientId)
	patient.PlannedYearDB.GetRowStruct(patientId)
	return patient
}

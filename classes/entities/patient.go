package entities

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
)

type Patient struct {
	PatientDB      *databaseworkers.PatientColumns
	PlanningYearDB []*databaseworkers.PlanningYearColumns
}

func PatientInit(patientId int, allPlans bool) *Patient {
	patient := &Patient{PatientDB: &databaseworkers.PatientColumns{}}
	patient.PatientDB.GetRowStruct(patientId)

	if !allPlans {
		var pyc *databaseworkers.PlanningYearColumns = &databaseworkers.PlanningYearColumns{}
		pyc.GetRowStruct(patientId)
		patient.PlanningYearDB = append(patient.PlanningYearDB, pyc)
	} else {
		patient.PlanningYearDB = databaseworkers.GetPatientPlans(patientId)
	}

	return patient
}

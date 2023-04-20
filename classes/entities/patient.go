package entities

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
)

type Patient struct {
	PatientDB    *databaseworkers.PatientColumns
	PlanningYear []*PlanningYear
}

func NewPatient(patientId int, allPlans bool) *Patient {
	patient := &Patient{PatientDB: &databaseworkers.PatientColumns{}}
	patient.PatientDB.GetRowStruct(patientId)

	if !allPlans {
		plYear := &PlanningYear{}
		plYear.NewPlanningYear(patientId)
		patient.PlanningYear = append(patient.PlanningYear, plYear)
	} else {
		var plYearDBStructs []*databaseworkers.PlanningYearColumns = databaseworkers.GetPatientPlans(patientId)
		for i := 0; i < len(plYearDBStructs); i++ {
			patient.PlanningYear = append(patient.PlanningYear, NewPlanningYear(plYearDBStructs[i]))
		}
	}

	return patient
}

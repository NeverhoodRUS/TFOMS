package entities

import (
	databaseworkers "tfoms_server/classes/dataBaseWorkers"
	"tfoms_server/static/strings"
)

type Patient struct {
	Id               int    // identificator
	ENP              string // enp
	SecondName       string
	FirstName        string
	MiddleName       string
	BirthDay         int // timestamp
	SexId            int
	Address          string
	PhoneNumber      string
	EndInsuranceDate int // timestamp
	OrganizationId   int // attached organization
	MedicalOrgId     int // medical organization id

	PriorityGroupId int
	PlYear          *PlanningYear
}

func PatientInit(patientId int) (*Patient, error) {
	row := databaseworkers.GetRow(strings.PatientTableName, patientId)
	patient := &Patient{}
	err := row.Scan(&patient)
	if err != nil {
		return nil, err
	}
	return patient, nil
}

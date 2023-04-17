package databaseworkers

import (
	"database/sql"
	"strconv"
	"tfoms_server/static/strings"
)

type PatientColumns struct {
	Id               int    `field:"id"`
	ENP              string `field:"enp"`
	FirstName        string `field:"firstname"`
	SecondName       string `field:"secondname"`
	MiddleName       string `field:"middlename"`
	BirthDate        int    `field:"birth_date"`
	Address          string `field:"address"`
	PhoneNumber      string `field:"phone_number"`
	EndInsuranceDate int    `field:"end_insurance_date"`
	InsuranceOrgId   int    `field:"insurance_org_id"`
	OrganizationId   int    `field:"organization_id"`
	SexId            int    `field:"sex_id"`
}

func GetPatientJournal(filters map[string]string) (*sql.Rows, error) {
	rows, err := getPostgresSession().Query(initQueryByFilter(filters))
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (pc *PatientColumns) GetRowStruct(id int) {
	var where string = "id = " + strconv.Itoa(id)
	getRow(strings.PatientTableName, where).Scan(&pc.Id, &pc.ENP, &pc.FirstName, &pc.SecondName, &pc.MiddleName,
		&pc.BirthDate, &pc.Address, &pc.PhoneNumber, &pc.EndInsuranceDate, &pc.InsuranceOrgId, &pc.OrganizationId, &pc.SexId)
}

func initQueryByFilter(filters map[string]string) string {
	if len(filters) == 0 {
		return "select id from " + strings.PatientTableName + " where id > 0"
	}
	var result string = "select id from " + strings.PatientTableName + " where "
	if val, ok := filters["organization_id"]; ok {
		result += "organization_id = " + val
	}
	if val, ok := filters["ensurance_org_id"]; ok {
		result += " and ensurance_org_id = " + val
	}
	if val, ok := filters["enp"]; ok {
		result += " and enp = " + val
	}
	if val, ok := filters["firstname"]; ok {
		result += " and firstname = " + val
	}
	if val, ok := filters["secondname"]; ok {
		result += " and secondname = " + val
	}
	if val, ok := filters["middlename"]; ok {
		result += " and middlename = " + val
	}
	if val, ok := filters["birth_date_before"]; ok {
		result += " and birthdate <= " + val
	}
	if val, ok := filters["birth_date_after"]; ok {
		result += " and birthdate >= " + val
	}
	return result
}

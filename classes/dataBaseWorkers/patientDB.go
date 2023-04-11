package databaseworkers

import (
	"database/sql"
	"tfoms_server/static/strings"
)

func GetPatientJournal(filters map[string]string) (*sql.Rows, error) {
	rows, err := db.Query(initQueryByFilter(filters))
	if err != nil {
		return nil, err
	}
	return rows, nil
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

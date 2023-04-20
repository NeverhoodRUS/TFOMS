package databaseworkers

import (
	"database/sql"
	"strconv"
	"strings"
	"tfoms_server/static/names"
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

func GetPatientJournal(filters map[string]string, limit int) (*sql.Rows, error) {
	rows, err := getPostgresSession().Query(initQueryByFilter(filters) + " limit " + strconv.Itoa(limit))
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (pc *PatientColumns) GetRowStruct(id int) {
	var where string = "id = " + strconv.Itoa(id)
	getRow(names.PatientTableName, where).Scan(&pc.Id, &pc.ENP, &pc.FirstName, &pc.SecondName, &pc.MiddleName,
		&pc.BirthDate, &pc.Address, &pc.PhoneNumber, &pc.EndInsuranceDate, &pc.InsuranceOrgId, &pc.OrganizationId, &pc.SexId)
}

func initQueryByFilter(filters map[string]string) string {
	//added where id > 0 just for using 'and' below without check
	var result string = "select id from " + names.PatientTableName + " as p where id > 0"
	if len(filters) == 0 {
		return result
	}

	setJoinsForFilter(&result, filters)
	setConditionsForFilter(&result, filters)

	return result
}

func setConditionsForFilter(result *string, filters map[string]string) {
	//NOT joined values
	if val, ok := filters["organization_id"]; ok {
		*result += " and organization_id = " + val
	}
	if val, ok := filters["ensurance_org_id"]; ok {
		*result += " and ensurance_org_id = " + val
	}
	if val, ok := filters["enp"]; ok {
		*result += " and enp = " + val
	}
	if val, ok := filters["firstname"]; ok {
		*result += " and firstname = " + val
	}
	if val, ok := filters["secondname"]; ok {
		*result += " and secondname = " + val
	}
	if val, ok := filters["middlename"]; ok {
		*result += " and middlename = " + val
	}
	if val, ok := filters["birth_date_before"]; ok {
		*result += " and birthdate <= " + val
	}
	if val, ok := filters["birth_date_after"]; ok {
		*result += " and birthdate >= " + val
	}
	//JOINED values
	if _, ok := filters["informing_date"]; ok {
		*result += " and inf.informing_date is not null"
	}
	if val, ok := filters["completed_event_id"]; ok {
		*result += " and pr.completed_event_id = " + val
	}
	if val, ok := filters["planned_year"]; ok {
		*result += " and ply.planned_year = " + val
	}
	if val, ok := filters["planned_month"]; ok {
		*result += " and ply.planned_month = " + val
	}
	if val, ok := filters["priority_group_id"]; ok {
		*result += "and ply.priority_group_id = " + val
	}
	if val, ok := filters["planned_event_id"]; ok {
		*result += "and ply.planned_event_id = " + val
	}
	if _, ok := filters["second_stage"]; ok {
		*result += " and pr.second_stage is not null"
	}
	if val, ok := filters["last_visit_date"]; ok {
		*result += "and ply.last_visit_date < " + val
	}
}

func setJoinsForFilter(result *string, filters map[string]string) {
	if _, ok := filters["informing_date"]; ok {
		*result += "join " + names.InformingTableName + " inf on inf.patient_id = p.id"
	}
	if _, ok := filters["completed_event_id"]; ok {
		*result += "join " + names.ProgressTableName + " pr on pr.patient_id = p.id"
	}
	if _, ok := filters["planned_year"]; ok {
		*result += "join " + names.PlanningYearTableName + " ply on ply.patient_id = p.id"
	}
	if _, ok := filters["planned_month"]; ok {
		if ok := strings.Contains(*result, names.PlanningYearTableName); !ok {
			*result += "join " + names.PlanningYearTableName + " ply on ply.patient_id = p.id"
		}
	}
	if _, ok := filters["priority_group_id"]; ok {
		if ok := strings.Contains(*result, names.PlanningYearTableName); !ok {
			*result += "join " + names.PlanningYearTableName + " ply on ply.patient_id = p.id"
		}
	}
	if _, ok := filters["planned_event_id"]; ok {
		if ok := strings.Contains(*result, names.PlanningYearTableName); !ok {
			*result += "join " + names.PlanningYearTableName + " ply on ply.patient_id = p.id"
		}
	}
	if _, ok := filters["second_stage"]; ok {
		if ok := strings.Contains(*result, names.ProgressTableName); !ok {
			*result += "join " + names.PlanningYearTableName + " pr on pr.patient_id = p.id"
		}
	}
	if _, ok := filters["last_visit_date"]; ok {
		if ok := strings.Contains(*result, names.PlanningYearTableName); !ok {
			*result += "join " + names.PlanningYearTableName + " ply on ply.patient_id = p.id"
		}
	}
}

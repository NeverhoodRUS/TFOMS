package databaseworkers

import (
	"strconv"
	"tfoms_server/static/names"
)

type PlanningYearColumns struct {
	Id              int `field:"id"`
	PlannedYear     int `field:"planned_year"`
	PlannedMonth    int `field:"planned_month"`
	LastVisitDate   int `field:"last_visit_date"`
	COVIDDate       int `field:"covid_date"`
	PatientId       int `field:"patient_id"`
	PlannedEventId  int `field:"planned_event_id"`
	PriorityGroupId int `field:"priority_group_id"`
}

func (pyc *PlanningYearColumns) GetRowStruct(patientId int) {
	var where string = "patient_id = " + strconv.Itoa(patientId) + " order by planned_year desc limit 1"
	getRow(names.PlanningYearTableName, where).Scan(&pyc.Id, &pyc.PlannedYear, &pyc.PlannedMonth, &pyc.LastVisitDate,
		&pyc.COVIDDate, &pyc.PatientId, &pyc.PlannedEventId, &pyc.PriorityGroupId)
}

func GetPatientPlans(patientId int) []*PlanningYearColumns {
	var where string = "patient_id = " + strconv.Itoa(patientId)
	rows, err := getRows(names.PlanningYearTableName, where)
	if err != nil {
		return nil
	}
	result := []*PlanningYearColumns{}
	for rows.Next() {
		var pyc *PlanningYearColumns
		err = rows.Scan(&pyc.Id, &pyc.PlannedYear, &pyc.PlannedMonth, &pyc.LastVisitDate, &pyc.COVIDDate,
			&pyc.PatientId, &pyc.PlannedEventId, &pyc.PriorityGroupId)
		if err == nil {
			return nil
		}
		result = append(result, pyc)
	}
	return result
}

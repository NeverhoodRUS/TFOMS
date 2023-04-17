package databaseworkers

import (
	"strconv"
	"tfoms_server/static/strings"
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
	var where string = "patient_id = " + strconv.Itoa(patientId)
	getRow(strings.PlanningYearTableName, where).Scan(&pyc.Id, &pyc.PlannedYear, &pyc.PlannedMonth, &pyc.LastVisitDate,
		&pyc.COVIDDate, &pyc.PatientId, &pyc.PlannedEventId, &pyc.PriorityGroupId)
}
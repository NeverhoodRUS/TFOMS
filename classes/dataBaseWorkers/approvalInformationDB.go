package databaseworkers

import (
	"strconv"
	"tfoms_server/static/names"
)

type ApprovalColumns struct {
	Id             int
	PlanningYearId int
	ApproveDate    int
	ApprovedUserId int
	InsuranceOrgId int
}

func (ac *ApprovalColumns) GetRowStruct(planningYearId int) {
	var where string = "planning_year_id = " + strconv.Itoa(planningYearId)
	getRow(names.PatientTableName, where).Scan(&ac.Id, &ac.PlanningYearId, &ac.ApproveDate, &ac.ApprovedUserId, &ac.InsuranceOrgId)
}

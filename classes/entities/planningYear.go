package entities

import databaseworkers "tfoms_server/classes/dataBaseWorkers"

type PlanningYear struct {
	PlanningYearDB      *databaseworkers.PlanningYearColumns
	Informing           *Informing
	ApprovalInformation *ApprovalInformation
	Progress            *Progress
}

func (plY *PlanningYear) NewPlanningYear(patientId int) {
	plY.PlanningYearDB.GetRowStruct(patientId)

	var plYearId int = plY.PlanningYearDB.Id

	plY.Informing.NewIforming(plYearId)
	plY.ApprovalInformation.NewApprovalInformation(plYearId)
	plY.Progress.NewProgress(plYearId)
}

func NewPlanningYear(planningYearDB *databaseworkers.PlanningYearColumns) *PlanningYear {
	result := &PlanningYear{}
	result.PlanningYearDB = planningYearDB

	var plYearId int = result.PlanningYearDB.Id

	result.Informing.NewIforming(plYearId)
	result.ApprovalInformation.NewApprovalInformation(plYearId)
	result.Progress.NewProgress(plYearId)

	return result
}

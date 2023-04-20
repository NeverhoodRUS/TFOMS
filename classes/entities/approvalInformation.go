package entities

import databaseworkers "tfoms_server/classes/dataBaseWorkers"

type ApprovalInformation struct {
	ApprovalDB *databaseworkers.ApprovalColumns
}

func (ai *ApprovalInformation) NewApprovalInformation(planningYearId int) {
	ai.ApprovalDB.GetRowStruct(planningYearId)
}

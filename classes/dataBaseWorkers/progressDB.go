package databaseworkers

import "tfoms_server/static/names"

type ProgressColumns struct {
	Id               int
	PlanningYearId   int
	EventDate        int
	CompletedEventId int
	HealthGroupId    int
	SecondStage      bool
	ExecutingOrgId   int
}

func (pc *ProgressColumns) GetRowStruct(planningYearId int) {
	var where string = ""
	getRow(names.ProgressTableName, where).Scan(&pc.Id, &pc.PlanningYearId, &pc.EventDate, &pc.CompletedEventId,
		&pc.HealthGroupId, &pc.SecondStage, &pc.ExecutingOrgId)
}

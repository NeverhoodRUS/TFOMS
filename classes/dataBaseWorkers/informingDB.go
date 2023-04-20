package databaseworkers

import (
	"strconv"
	"tfoms_server/static/names"
)

type InformingColumns struct {
	Id                int
	PlanningYearId    int
	InformingDate     int
	InformingTypeId   int
	InformingMethodId int
}

func (ic *InformingColumns) GetRowStruct(planningYearId int) {
	var where string = "planning_year_id = " + strconv.Itoa(planningYearId)
	getRow(names.InformingTableName, where).Scan(&ic.Id, &ic.PlanningYearId, &ic.InformingDate, &ic.InformingTypeId,
		&ic.InformingMethodId)
}

package entities

import databaseworkers "tfoms_server/classes/dataBaseWorkers"

type Progress struct {
	ProgressDB *databaseworkers.ProgressColumns
}

func (p *Progress) NewProgress(planningYearId int) {
	p.ProgressDB.GetRowStruct(planningYearId)
}

package entities

import databaseworkers "tfoms_server/classes/dataBaseWorkers"

type Informing struct {
	InformingDB *databaseworkers.InformingColumns
}

func (i *Informing) NewIforming(planningYearId int) {
	i.InformingDB.GetRowStruct(planningYearId)
}

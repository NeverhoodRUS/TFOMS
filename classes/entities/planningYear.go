package entities

type PlanningYear struct {
	Id              int
	Year            int
	PatientId       int
	LastVisitDate   int
	COVIDDate       int
	PlannedEvent    int
	PlannedMonth    int
	PriorityGroupId int
}

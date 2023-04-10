package entities

type Progress struct {
	Id                    int
	PatientId             int
	EventDate             int
	CompletedEventId      int
	HealthGroupId         int
	RefferalSecondStage   bool
	ExecutingOrganization int
}

package entities

type Patient struct {
	Id               int
	ENP              string
	MedicalOrgId     int
	SecondName       string
	FirstName        string
	MiddleName       string
	BirthDay         int
	SexId            int
	OrganizationId   int
	Address          string
	PhoneNumber      string
	EndEnsuranceDate int
}

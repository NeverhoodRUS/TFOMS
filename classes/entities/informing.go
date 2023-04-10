package entities

type Informing struct {
	Id                int
	PatientId         int
	InformingDate     int
	InformingTypeId   int //Вид информирования
	InformingMethodId int //Способ информирования
}

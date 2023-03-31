package entities

type User struct {
	id       int
	roleId   int
	session  string
	objectId int
}

func NewUser(id int, roleId int, objectId int) *User {
	return &User{id, roleId, "", objectId}
}

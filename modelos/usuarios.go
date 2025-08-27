package modelos

import "time"

type User struct {
	userId    int
	Nombre    string
	CreatedAt time.Time
	Status    bool
}

func (usuario *User) AddUser(userid int, name string, createdAt time.Time, status bool) {

	usuario.userId = userid
	usuario.Nombre = name
	usuario.CreatedAt = createdAt
	usuario.Status = status
}

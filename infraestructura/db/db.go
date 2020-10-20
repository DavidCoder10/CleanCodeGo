package db

import (
	"CleanCode/entidades_negocio"
)

//Database ...
type Database struct {
	Users map[int]entidades_negocio.EUser
}

//NewDB ...
func NewDB() Idb {
	return &Database{}
}

//Data ...
func (Database Database) Data() map[int]entidades_negocio.EUser {
	return map[int]entidades_negocio.EUser{
		0: entidades_negocio.EUser{
			ID:    1,
			Name:  "David",
			Email: "David@",
			Age:   18,
		},
	}
}

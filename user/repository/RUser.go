package repository

import (
	"CleanCode/infraestructura/db"
	"CleanCode/entidades_negocio"
	"CleanCode/user/contracts"
	"errors"
)

//RUser ...
type RUser struct {
	db db.Idb
}

//NewRUserRepo ...
func NewRUserRepo(idb db.Idb) contracts.IRepository {
	return &RUser{db: idb}
}

//GetUserByID ...
func (rUser *RUser) GetUserByID(id int) (entidades_negocio.EUser, error) {

	data := rUser.db.Data()
	user, found := data[id]

	if found {
		return user, nil
	}

	return user, errors.New("Register not found")
}

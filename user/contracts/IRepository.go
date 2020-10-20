package contracts

import "CleanCode/entidades_negocio"

//IRepository ...
type IRepository interface {
	GetUserByID(id int) (entidades_negocio.EUser, error)
}

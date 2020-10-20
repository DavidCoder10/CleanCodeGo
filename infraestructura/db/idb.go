package db

import "CleanCode/entidades_negocio"

//Idb ...
type Idb interface {
	Data() map[int]entidades_negocio.EUser
}

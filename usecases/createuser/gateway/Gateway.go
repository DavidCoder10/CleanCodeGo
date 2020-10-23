package gateway

import (
	"CleanCodeGo/entities/user"
)

type Gateway interface {
	FindByEmail(email string) (user.User, error)
	Create(user user.User) (user.User, error)
}

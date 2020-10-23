package port

import "CleanCodeGo/user/domain/entity"

type UserRepository interface {
	FindByEmail(email string) (entity.User, error)
	Create(user entity.User) (entity.User, error)
}
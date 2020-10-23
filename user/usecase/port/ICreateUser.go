package port

import (
	"CleanCodeGo/user/usecase/request"
	"CleanCodeGo/user/usecase/response"
)

type ICreateUser interface {
	Create(user request.UserRequest) (response.UserResponse, error)
}

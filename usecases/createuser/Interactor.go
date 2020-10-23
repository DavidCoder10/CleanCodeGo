package createuser

import (
	"CleanCodeGo/usecases/createuser/request"
	"CleanCodeGo/usecases/createuser/response"
)

type Interactor interface {
	Create(user request.UserRequest) (response.UserResponse, error)
}

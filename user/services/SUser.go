package services

import (
	"CleanCode/user/contracts"
	"errors"
	"fmt"
)

//SUser ...
type SUser struct {
	rUser contracts.IRepository
}

//NewSUserService ...
func NewSUserService(rUseri contracts.IRepository) contracts.IService {
	return &SUser{rUser: rUseri}
}

//ValidateUserOfLegalAge ...
func (sUser *SUser) ValidateUserOfLegalAge(id int) (string, error) {

	user, err := sUser.rUser.GetUserByID(id)

	if err != nil {
		return "Error", err
	}

	if user.Age < 18 {
		return "Error", errors.New("The user is not of legal age")
	}

	return fmt.Sprintf("Welcome %s", user.Name), nil
}

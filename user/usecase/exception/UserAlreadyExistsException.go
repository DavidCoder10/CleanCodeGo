package exception

import "errors"

type UserAlreadyExistsException struct {}

func (userAlreadyExistsException *UserAlreadyExistsException) UserAlreadyExistsException(message string) error {
	return errors.New(message)
}
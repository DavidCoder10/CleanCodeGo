package main

import (
	"CleanCodeGo/user/adapter/Repository/DB/Sql/mysql"
	"CleanCodeGo/user/adapter/controller"
	"CleanCodeGo/user/adapter/controller/request"
	"CleanCodeGo/user/usecase"
	"fmt"
	"github.com/jinzhu/gorm"
)

func main() {

	repo := mysql.NewUserRepository(&gorm.DB{})
	useCase := usecase.NewCreateUser(repo)
	controller := controller.NewUserController(useCase)

	request := request.UserRequest{Name: "David", Email: "David@", Password: "123"}

	response, err := controller.CreateUser(request)

	if err != nil {
		panic(err)
	}

	fmt.Sprint(response)
}

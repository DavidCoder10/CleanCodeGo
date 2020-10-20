package main

import (
	"CleanCode/infraestructura/db"
	"CleanCode/user/repository"
	"CleanCode/user/services"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	db := db.NewDB()
	r := repository.NewRUserRepo(db)
	s := services.NewSUserService(r)
	user, err := s.ValidateUserOfLegalAge(0)

	if err != nil {
		panic(color.RedString(err.Error()))
	}

	fmt.Printf(color.HiGreenString(user))
}

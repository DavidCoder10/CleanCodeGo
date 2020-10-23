package mysql

import (
	"CleanCodeGo/entities/user"
	"CleanCodeGo/usecases/createuser/gateway"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) gateway.Gateway {

	connection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"dbroot",
		"127.0.0.1",
		"3307",
		"users",
	)

	database, err := gorm.Open("mysql", connection)

	if err != nil {
		panic(fmt.Sprint("Connection refused! %s", err))
	}
	fmt.Print("Connection success")
	return &UserRepository{db: database}
}

func (userRepository *UserRepository) FindByEmail(email string) (user.User, error) {
	user := user.User{Email: email}
	err := userRepository.db.Where(&user).First(&user).Error
	return user, err
}

func (userRepository *UserRepository) Create(user user.User) (user.User, error) {
	err := userRepository.db.Create(user).Error
	return user, err
}

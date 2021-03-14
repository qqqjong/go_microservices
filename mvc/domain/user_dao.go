package domain

import (
	"fmt"
	"github.com/qqqjong/go_microservices/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: &User{Id: 123, FirstName: "Jong", LastName: "Kim", Email: "aaa@aaa.com"},
	}

	UserDao userServiceInterface
)

func init() {
	UserDao = &userDao{}
}

type userServiceInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v does not exists", userId),
		StatusCode: http.StatusNotFound,
		Code: "not_found",
	}

}
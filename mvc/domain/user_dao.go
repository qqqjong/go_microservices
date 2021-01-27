package domain

import (
	"fmt"
	"github.com/qqqjong/go_microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: &User{Id: 123, FirstName: "Jong", LastName: "Kim", Email: "aaa@aaa.com"},
	}

	UserDao userDao
)

type userDao struct {

}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code: "not_found",
	}

}
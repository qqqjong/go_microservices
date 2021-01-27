package services

import (
	"github.com/qqqjong/go_microservices/mvc/domain"
	"github.com/qqqjong/go_microservices/mvc/utils"
)

type userService struct {

}

var (
	UsersService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
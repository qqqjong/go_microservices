package services

import (
	"github.com/qqqjong/go_microservices/mvc/domain"
	"github.com/qqqjong/go_microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
package services

import (
	"github.com/qqqjong/go_microservices/mvc/domain"
	"github.com/qqqjong/go_microservices/mvc/utils"
	"net/http"
)

type itemsService struct {

}

var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message: "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
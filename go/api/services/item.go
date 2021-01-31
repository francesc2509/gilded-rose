package services

import (
	"github.com/francesc2509/glided-rose/core/repositories"
	"github.com/francesc2509/glided-rose/entities"
)

type itemService struct{}

func (service *itemService) Get() ([]*entities.Item, error) {
	return repositories.Item.Get()
}

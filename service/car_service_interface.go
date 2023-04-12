package service

import "github.com/aabdullahgungor/go-restapi-mock/model"

type ICarService interface {
	GetAll() ([]model.Car, error)
	GetById(id string) (model.Car, error)
	Create(car *model.Car) error
	Edit(car *model.Car) error
	Delete(id string) error
}
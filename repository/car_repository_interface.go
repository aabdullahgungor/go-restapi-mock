package repository

import "github.com/aabdullahgungor/go-restapi-mock/model"

type ICarRepository interface {
	GetAllCars() ([]model.Car, error)
	GetCarById(id string) (model.Car, error)
	CreateCar(car *model.Car) error
	EditCar(car *model.Car) error
	DeleteCar(id string) error
}
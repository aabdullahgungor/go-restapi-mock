package service

import (
	"testing"
	"time"

	"github.com/aabdullahgungor/go-restapi-mock/model"
	"github.com/aabdullahgungor/go-restapi-mock/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func TestDefaultCarService_GetAll(t *testing.T) {

	mockCtrl := gomock.NewController(t)
  	defer mockCtrl.Finish()
	mockRepository := repository.NewMockICarRepository(mockCtrl)
	mockRepository.EXPECT().GetAllCars().Return([]model.Car{{
		Id : primitive.ObjectID{'1'},
    	Brand : "bmw",  
    	Series : "",  
    	Year   : time.Time{}, 
   		Fuel   : "diesel",  
    	Gear   :"manuel", 
    	Situation : "secondhand",
    	Km        : 100000,
   		Color     : "white",
    	Price     : 10000, },
		{
		Id : primitive.ObjectID{'2'},
    	Brand : "fiat",  
    	Series : "",  
    	Year   : time.Time{}, 
   		Fuel   : "petrol",  
    	Gear   :"manuel", 
    	Situation : "secondhand",
    	Km        : 120000,
   		Color     : "black",
    	Price     : 10000,
	}}, nil)

	carService := NewDefaultCarService(mockRepository)
	cars , err := carService.GetAll()

	if assert.Nil(t,err) {
		if len(cars) == 2 {
			t.Log("Car counts is matching, func run succesfuly")
		} else {
			t.Log("Car counts not matching, there is a problem in func")
		}
	}else {
		t.Log(err)
	}
}

func TestDefaultCarService_GetById(t *testing.T) {
	id := "1"
	/*mockCarResp := model.Car{
		Id : primitive.ObjectID{'1'},
    	Brand : "bmw",  
    	Series : "",  
    	Year   : time.Time{}, 
   		Fuel   : "diesel",  
    	Gear   :"manuel", 
    	Situation : "secondhand",
    	Km        : 100000,
   		Color     : "white",
    	Price     : 10000,
	}*/
	mockCtrl := gomock.NewController(t)
  	defer mockCtrl.Finish()
	mockRepository := repository.NewMockICarRepository(mockCtrl)
	mockRepository.EXPECT().GetCarById(gomock.Eq(id)).Return(model.Car{}, ErrCarNotFound)
	
	carService := NewDefaultCarService(mockRepository)
	_ , err := carService.GetById(id)

	assert.ErrorIs(t, err, ErrCarNotFound)
	// if car.Brand != "fiat" {
	// 	t.Errorf("Brand is not maching")
	// }
}

func TestDefaultCarService_Create(t *testing.T) {
	
	car := model.Car{Brand: "Bmw"}
	mockCtrl := gomock.NewController(t)
  	defer mockCtrl.Finish()
	mockRepository := repository.NewMockICarRepository(mockCtrl)
	mockRepository.EXPECT().CreateCar(&car).Return(nil).Times(1)

	carService := NewDefaultCarService(mockRepository)
	err := carService.Create(&car)

	if assert.Nil(t, err) {
		t.Log("Success Create Car") 
	} else {
		t.Log("Car cannot create")
	}
}

func TestDefaultCarService_Edit(t *testing.T) {

	car := model.Car{Brand: "Bmw"}
	mockCtrl := gomock.NewController(t)
  	defer mockCtrl.Finish()
	mockRepository := repository.NewMockICarRepository(mockCtrl)
	mockRepository.EXPECT().EditCar(&car).Return(nil).Times(1)

	carService := NewDefaultCarService(mockRepository)
	err := carService.Edit(&car)

	if assert.Nil(t, err) {
		t.Log("Success Update Car") 
	} else {
		t.Log(err)
	}
}

func TestDefaultCarService_Delete(t *testing.T) {

	id := "1"
	mockCtrl := gomock.NewController(t)
  	defer mockCtrl.Finish()
	mockRepository := repository.NewMockICarRepository(mockCtrl)
	mockRepository.EXPECT().DeleteCar(gomock.Eq(id)).Return(nil).Times(1)
	
	carService := NewDefaultCarService(mockRepository)
	err := carService.Delete(id)

	if assert.Nil(t, err) {
		t.Log("Success delete Car") 
	} else {
		t.Log(err)
	}
}


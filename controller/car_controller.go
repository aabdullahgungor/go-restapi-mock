package controller

import (
	"errors"
	"net/http"

	"github.com/aabdullahgungor/go-restapi-mock/model"
	"github.com/aabdullahgungor/go-restapi-mock/service"
	"github.com/gin-gonic/gin"
)

type carController struct {
	service service.ICarService
}

func NewCarController(cs service.ICarService) *carController {
	return &carController{service: cs}
}

func (cs *carController) GetAllCars(c *gin.Context) {
	cars, err := cs.service.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound,gin.H{"error": "Cars cannot show: " + err.Error(), })
		return
	}
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, cars)
}

func (cs *carController) GetCarById(c *gin.Context) { 
	str_id := c.Param("id")
	car, err := cs.service.GetById(str_id)
	if err != nil {
	 	if errors.Is(err, service.ErrIDIsNotValid) {
	 		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	 		return
	 	} else if  errors.Is(err, service.ErrCarNotFound) {
	 		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	 		return
	 	}
	 	c.IndentedJSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	 	return
	} 
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, car)
}

func (cs *carController) CreateCar(c *gin.Context) { 
	car := model.Car{}
	err := c.BindJSON(&car)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = cs.service.Create(&car)

	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot create car: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message":"Car has been created"})
}		

func (cs *carController) EditCar(c *gin.Context) { 
	var car model.Car
	err := c.ShouldBindJSON(&car)

	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = cs.service.Edit(&car)

	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot edit car: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message":"Car has been edited","car_id": car.Id})
}		

func (cs *carController) DeleteCar(c *gin.Context) { 
	str_id := c.Param("id")
	err := cs.service.Delete(str_id)
	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "id is not valid"+err.Error()})
			return
		} else if  errors.Is(err, service.ErrCarNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Car cannot be found"+err.Error()})
			return
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message":"Car has been deleted","car_id": str_id})	
}		
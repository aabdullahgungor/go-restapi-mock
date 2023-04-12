package router

import (
	"github.com/aabdullahgungor/go-restapi-mock/controller"
	"github.com/aabdullahgungor/go-restapi-mock/repository"
	"github.com/aabdullahgungor/go-restapi-mock/service"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	carRepo := repository.NewMongoDbCarRepository()
	carService := service.NewDefaultCarService(carRepo)
	carController := controller.NewCarController(carService)

	main := router.Group("api/v1")
	{
		cars := main.Group("cars")
		{
			cars.GET("/", carController.GetAllCars)
			cars.GET("/:id", carController.GetCarById)
			cars.POST("/", carController.CreateCar)
			cars.PUT("/", carController.EditCar)
			cars.DELETE("/:id", carController.DeleteCar)
		}
	}
	return router
}
package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/aabdullahgungor/go-restapi-mock/model"
	"github.com/aabdullahgungor/go-restapi-mock/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCarController_GetAllCars(t *testing.T) {

	t.Run("Error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockICarService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]model.Car{}, errors.New("hata!")).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		carTestController := NewCarController(mockService)
		carTestController.GetAllCars(ctx)

		req, _ := http.NewRequest("GET", "api/v1/cars", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockICarService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]model.Car{}, nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		carTestController := NewCarController(mockService)
		carTestController.GetAllCars(ctx)

		req, _ := http.NewRequest("GET", "api/v1/cars", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println(w.Code)
	})
}

func TestCarController_GetCarById(t *testing.T) {

	t.Run("Error", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockICarService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(model.Car{}, service.ErrCarNotFound).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		carTestController := NewCarController(mockService)
		carTestController.GetCarById(ctx)

		req, _ := http.NewRequest("GET", "api/v1/cars/:id", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockICarService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(model.Car{
			Brand:     "bmw",
			Series:    "",
			Year:      time.Time{},
			Fuel:      "diesel",
			Gear:      "manuel",
			Situation: "secondhand",
			Km:        100000,
			Color:     "white",
			Price:     10000,
		}, nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		carTestController := NewCarController(mockService)
		carTestController.GetCarById(ctx)

		req, _ := http.NewRequest("GET", "api/v1/cars/:id", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println(w.Code)

		var responseData model.Car
		json.NewDecoder(w.Body).Decode(&responseData)
		assert.Equal(t, "bmw", responseData.Brand)
		t.Log("\nCar brand is: ", responseData.Brand)
	})
}

func TestCarController_CreateCar(t *testing.T) {

	t.Run("ErrorCreate", func(t *testing.T) {
		car := model.Car{Brand: "Bmw"}
		jsonValue, _ := json.Marshal(car)
		byteCar := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockICarService(mockCtrl)
		mockService.EXPECT().Create(&car).Return(errors.New("hata")).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		carTestController := NewCarController(mockService)
		carTestController.CreateCar(ctx)
		req, err := http.NewRequest("POST", "api/v1/cars", byteCar)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		car := model.Car{Brand: "Bmw"}
		jsonValue, _ := json.Marshal(car)
		byteCar := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockICarService(mockCtrl)
		mockService.EXPECT().Create(&car).Return(nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		carTestController := NewCarController(mockService)
		carTestController.CreateCar(ctx)
		req, err := http.NewRequest("POST", "api/v1/cars", byteCar)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		t.Log(w.Body.String())
	})
}

func TestCarController_EditCar(t *testing.T) {
	t.Run("ErrorEdit", func(t *testing.T) {
		car := model.Car{Brand: "Bmw"}
		jsonValue, _ := json.Marshal(car)
		byteCar := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockICarService(mockCtrl)
		mockService.EXPECT().Edit(&car).Return(errors.New("hata")).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		carTestController := NewCarController(mockService)
		carTestController.EditCar(ctx)
		req, err := http.NewRequest("PUT", "api/v1/cars", byteCar)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotAcceptable, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		car := model.Car{Brand: "Bmw"}
		jsonValue, _ := json.Marshal(car)
		byteCar := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockICarService(mockCtrl)
		mockService.EXPECT().Edit(&car).Return(nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		carTestController := NewCarController(mockService)
		carTestController.EditCar(ctx)
		req, err := http.NewRequest("PUT", "api/v1/cars", byteCar)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		t.Log(w.Body.String())
	})
}

func TestCarController_DeleteCar(t *testing.T) {

	t.Run("Error", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockICarService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(service.ErrCarNotFound).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, _ := gin.CreateTestContext(w)
		carTestController := NewCarController(mockService)
		carTestController.DeleteCar(ctx)
		assert.Equal(t, http.StatusNotFound, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockICarService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, _ := gin.CreateTestContext(w)
		carTestController := NewCarController(mockService)
		carTestController.DeleteCar(ctx)
		assert.Equal(t, http.StatusAccepted, w.Code)
		t.Log(w.Body.String())
	})

}

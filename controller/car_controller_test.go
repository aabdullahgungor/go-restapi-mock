package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
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

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		carTestController := NewCarController(mockService)
		carTestController.GetAllCars(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockICarService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]model.Car{}, nil).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		carTestController := NewCarController(mockService)
		carTestController.GetAllCars(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println(w.Code)
	})
}

func TestCarController_GetCarById(t *testing.T) {

	t.Run("Error", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockICarService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(model.Car{}, service.ErrCarNotFound).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		params := []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}
		ctx.Params = params
		carTestController := NewCarController(mockService)
		carTestController.GetCarById(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		id := "1"
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

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		params := []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}
		ctx.Params = params
		carTestController := NewCarController(mockService)
		carTestController.GetCarById(ctx)

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

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(byteCar)
		carTestController := NewCarController(mockService)
		carTestController.CreateCar(ctx)

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
		mockService.EXPECT().Create(&car).Return(nil).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(byteCar)
		carTestController := NewCarController(mockService)
		carTestController.CreateCar(ctx)

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

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(byteCar)
		carTestController := NewCarController(mockService)
		carTestController.EditCar(ctx)

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

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "PUT"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(byteCar)
		carTestController := NewCarController(mockService)
		carTestController.EditCar(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		t.Log(w.Body.String())
	})
}

func TestCarController_DeleteCar(t *testing.T) {

	t.Run("Error", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockICarService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(service.ErrCarNotFound).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "DELETE"
		ctx.Request.Header.Set("Content-Type", "application/json")
		params := []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}
		ctx.Params = params
		carTestController := NewCarController(mockService)
		carTestController.DeleteCar(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := service.NewMockICarService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(nil).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "DELETE"
		ctx.Request.Header.Set("Content-Type", "application/json")
		params := []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}
		ctx.Params = params
		carTestController := NewCarController(mockService)
		carTestController.DeleteCar(ctx)

		assert.Equal(t, http.StatusAccepted, w.Code)
		t.Log(w.Body.String())
	})

}

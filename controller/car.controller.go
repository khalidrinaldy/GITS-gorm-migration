package controller

import (
	"migration/entity"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetCars(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cars []entity.Car
		result := db.Find(&cars)
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, cars)
	}
}

func AddCar(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var car entity.Car
		car.ModelName= c.FormValue("modelName")

		result := db.Create(&car)
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, car)
	}
}

func UpdateCar(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var car entity.Car
		car.ModelName= c.FormValue("modelName")
		
		result := db.Model(&car).Where("id = ?", c.Param("id")).Update("name", car.ModelName)
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, car)
	}
}

func DeleteCar(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var car entity.Car
		result := db.Delete(&car, c.Param("id"))
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, car)
	}
}
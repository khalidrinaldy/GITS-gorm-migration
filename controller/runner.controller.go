package controller

import (
	"migration/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetRunners(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var runners []entity.Runner
		result := db.Find(&runners)
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, runners)
	}
}

func AddRunner(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var runner entity.Runner
		runner.Speed, _ = strconv.Atoi(c.FormValue("speed"))

		result := db.Create(&runner)
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, runner)
	}
}

func UpdateRunner(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var runner entity.Runner
		runner.Speed, _ = strconv.Atoi(c.FormValue("speed"))
		
		result := db.Model(&runner).Where("id = ?", c.Param("id")).Update("name", runner.Speed)
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, runner)
	}
}

func DeleteRunner(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var runner entity.Runner
		result := db.Delete(&runner, c.Param("id"))
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, runner)
	}
}
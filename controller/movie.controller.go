package controller

import (
	"migration/entity"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetMovies(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var movies []entity.Movie
		result := db.Find(&movies)
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, movies)
	}
}

func AddMovie(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var movie entity.Movie
		movie.Title = c.FormValue("title")

		result := db.Create(&movie)
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, movie)
	}
}

func UpdateMovie(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var movie entity.Movie
		movie.Title = c.FormValue("title")
		
		result := db.Model(&movie).Where("id = ?", c.Param("id")).Update("name", movie.Title)
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, movie)
	}
}

func DeleteMovie(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var movie entity.Movie
		result := db.Delete(&movie, c.Param("id"))
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, movie)
	}
}
package controller

import (
	"migration/entity"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetBooks(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var books []entity.Book
		result := db.Find(&books)
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, books)
	}
}

func AddBook(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var book entity.Book
		book.Title = c.FormValue("title")

		result := db.Create(&book)
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, book)
	}
}

func UpdateBook(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var book entity.Book
		book.Title = c.FormValue("title")
		
		result := db.Model(&book).Where("id = ?", c.Param("id")).Update("name", book.Title)
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, book)
	}
}

func DeleteBook(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var book entity.Book
		result := db.Delete(&book, c.Param("id"))
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, book)
	}
}
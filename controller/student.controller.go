package controller

import (
	"migration/entity"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetStudents(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var students []entity.Student
		result := db.Find(&students)
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, students)
	}
}

func AddStudent(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var student entity.Student
		student.Name = c.FormValue("name")

		result := db.Create(&student)
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, student)
	}
}

func UpdateStudent(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var student entity.Student
		student.Name = c.FormValue("name")
		
		result := db.Model(&student).Where("id = ?", c.Param("id")).Update("name", student.Name)
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, student)
	}
}

func DeleteStudent(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var student entity.Student
		result := db.Delete(&student, c.Param("id"))
		if result.Error!=nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, student)
	}
}
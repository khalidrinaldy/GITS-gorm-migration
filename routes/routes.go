package routes

import (
	"fmt"
	"migration/config"
	"migration/entity"
	"migration/controller"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {
	//init echo
	ech := echo.New()

	//Config
	cfg, err := config.NewConfig(".env")
	if err != nil {
		panic(err)
	}

	//Database
	db := openDatabase(cfg)
	db.AutoMigrate(&entity.Student{})

	//Student Routes
	ech.GET("/students", controller.GetStudents(db))
	ech.POST("/students", controller.AddStudent(db))
	ech.PUT("/students/:id", controller.UpdateStudent(db))
	ech.DELETE("/students/:id", controller.DeleteStudent(db))

	//Book Routes
	ech.GET("/books", controller.GetBooks(db))
	ech.POST("/books", controller.AddBook(db))
	ech.PUT("/books/:id", controller.UpdateBook(db))
	ech.DELETE("/books/:id", controller.DeleteBook(db))

	//Movie Routes
	ech.GET("/movies", controller.GetMovies(db))
	ech.POST("/movies", controller.AddMovie(db))
	ech.PUT("/movies/:id", controller.UpdateMovie(db))
	ech.DELETE("/movies/:id", controller.DeleteMovie(db))

	//Runner Routes
	ech.GET("/runners", controller.GetRunners(db))
	ech.POST("/runners", controller.AddRunner(db))
	ech.PUT("/runners/:id", controller.UpdateRunner(db))
	ech.DELETE("/runners/:id", controller.DeleteRunner(db))

	//Car Routes
	ech.GET("/cars", controller.GetCars(db))
	ech.POST("/cars", controller.AddCar(db))
	ech.PUT("/cars/:id", controller.UpdateCar(db))
	ech.DELETE("/cars/:id", controller.DeleteCar(db))

	return ech
}

func openDatabase(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name)
	db,err := gorm.Open(postgres.Open(string(dsn)), &gorm.Config{})
	checkError(err)
	return db
}

func checkError(err error)  {
	if err!=nil {
		panic(err)
	}
}
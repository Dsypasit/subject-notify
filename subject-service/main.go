package main

import (
	"subject-service/api"
	"subject-service/repositories"
	"subject-service/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var subject repositories.Subject

func main() {
	db := initialDatabase()
	subjectRepo := repositories.NewSubjectRepository(db)
	subjectService := services.NewSubjectService(subjectRepo)
	subjectHandler := api.NewSubjectHandler(subjectService)

	app := fiber.New()
	api.Route(app, subjectHandler)
	app.Listen("127.0.0.1:800")

}

func initialDatabase() *gorm.DB {
	dsn := "root:1234@tcp(localhost:4307)/users"
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	return db
}

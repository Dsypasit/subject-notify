package main

import (
	"user-service/api"
	"user-service/repositories"
	"user-service/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db := initialDatabase()
	userRepo := repositories.NewUserRepostoryDB(db)
	userService := services.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)

	app := fiber.New()
	api.Route(app, userHandler)
	app.Listen("127.0.0.1:8000")
}

func initialDatabase() *gorm.DB {
	dsn := "root:1234@tcp(localhost:4306)/users"
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	return db
}

package main

import (
	"fmt"
	"user-service/repositories"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := "root:1234@tcp(localhost:4306)/users"
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	userRepo := repositories.NewUserRepostoryDB(db)
	users, err := userRepo.GetUserByUsername("hi5")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", users)
}

package main

import (
	"fmt"
	"user-service/repositories"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:1234@tcp(localhost:4306)/users"
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	userRepo := repositories.NewUserRepostoryDB(db)
	user := repositories.User{
		Username: "hi4",
		Password: "1234",
	}
	userRepo.Create(user)
	// userUpdate := repositories.UpdateUserInformation{
	// 	Username: "hi2",
	// 	LineID:   "kkk",
	// }
	// userRepo.UpdateInformatin(userUpdate)
	users, err := userRepo.GetUserByUsername("hi4")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", users)
}

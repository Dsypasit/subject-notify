package main

import (
	"fmt"
	"subject-service/repositories"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var subject repositories.Subject

func main() {
	db := initialDatabase()
	subjectRepo := repositories.NewSubjectRepository(db)

	subject = repositories.Subject{
		Username:    "Pasit",
		SubjectName: "English",
		Day:         "Mon",
		Time:        "Afternoon",
		Teacher:     "Somsak",
		Link:        "google.com",
		Active:      true,
	}

	subQuery := repositories.SubjectQuery{
		SubjectName: "Math",
	}

	// subjectRepo.CreateSubject(subject)
	subs, err := subjectRepo.GetSubjectByQuery(subQuery)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", subs)
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

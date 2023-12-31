package main

import (
	"gorm.io/gorm"

	"gorm.io/driver/postgres"

	"fmt"

	"net/http"
)

var db *gorm.DB
var err error

type Number struct {
	gorm.Model
	Number int
	IsOdd  bool
}

func ConnectDb(POSTGRES_PASSWORD string) {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		"postgres",
		POSTGRES_PASSWORD,
		"numbers",
	)
	db, _ = gorm.Open(postgres.Open(dsn))
}

func CreateTable() {
	db.AutoMigrate(&Number{})

	for i := 0; i <= 1000; i++ {
		InsertNumber(i, i%2 == 1)
	}
}

func InsertNumber(number int, isOdd bool) {
	insertValue := Number{Number: number, IsOdd: isOdd}
	db.Create(&insertValue)
}

func GetNumber(number int) (Number, int) {
	var statusCode int
	selectValue := Number{Number: number}
	result := db.Where("number = ?", number).Find(&selectValue)

	if result.RowsAffected == 1 {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusNotFound
	}

	return selectValue, statusCode
}

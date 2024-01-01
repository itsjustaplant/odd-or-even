package main

import (
	"gorm.io/gorm"

	"gorm.io/driver/postgres"

	"fmt"

	"net/http"

	"math"
)

var db *gorm.DB

// var err error -> need error handling

type Number struct {
	gorm.Model
	Number int
	IsOdd  bool
}

func ConnectDb(POSTGRES_PASSWORD string) {
	dsn := fmt.Sprintf(
		"host=db user=postgres password=%s dbname=numbers port=5432 sslmode=disable",
		POSTGRES_PASSWORD,
	)
	db, _ = gorm.Open(postgres.Open(dsn))
}

func CreateTable() {
	if !db.Migrator().HasTable("numbers") {
		fmt.Println("Creating numbers table and inserting values")
		db.AutoMigrate(&Number{})

		for i := 0; i < math.MaxInt16; i++ {
			InsertNumber(i, i%2 == 1)
		}
		fmt.Println("Values are inserted and db is ready")
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

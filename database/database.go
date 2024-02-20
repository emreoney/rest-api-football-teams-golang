package database

import (
	"gomod/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const dbConnection = "host=localhost port=5432 user=postgres dbname=footballProject password=Alexabi12312 sslmode=disable"

func Init() {
	DB, err = gorm.Open(postgres.Open(dbConnection), &gorm.Config{})

	DB.AutoMigrate(models.League{}, models.Player{}, models.Team{}, models.Account{})
}

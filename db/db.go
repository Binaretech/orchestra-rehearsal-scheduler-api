package db

import (
	"fmt"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect connects to the database
func Connect() (*gorm.DB, error) {
	cnf := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", cnf.DatabseHost, cnf.DatabaseUser, cnf.DatabasePass, cnf.DatabaseName, cnf.DatabasePort)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return DB, err
}

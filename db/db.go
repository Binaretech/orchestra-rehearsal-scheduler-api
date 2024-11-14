package db

import (
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	cnf := config.GetConfig()

	dsn := cnf.DatabaseURL

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return DB, err
}

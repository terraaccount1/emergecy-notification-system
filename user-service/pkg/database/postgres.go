package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect db : err : %+v", err))
	}

	sqldb, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("failed to get sqldb : err : %+v", err))
	}

	err = sqldb.Ping()
	if err != nil {
		panic(fmt.Sprintf("failed to ping db : err : %+v", err))
	}

	return db
}

func MigrateModels(db *gorm.DB, model interface{}) {
	err := db.Debug().AutoMigrate(model)
	if err != nil {
		panic(fmt.Sprintf("failed to auto migrate db : err : %+v", err))
	}
}

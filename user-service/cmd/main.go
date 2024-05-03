package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/terraaccount1/emergecy-notification-system/user-service/internal/models"
	"github.com/terraaccount1/emergecy-notification-system/user-service/pkg/database"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func main() {

}

func init() {
	godotenv.Load("../.env")
	dsn := fmt.Sprint(
		"host=", os.Getenv("PG_HOST"),
		" user=", os.Getenv("PG_USER"),
		" password=", os.Getenv("PG_PASS"),
		" dbname=", os.Getenv("PG_DB"),
		" port=", os.Getenv("PG_PORT"),
		" sslmode=disable",
	)
	db = database.NewPostgresConnection(dsn)
	database.MigrateModels(db, &models.Contact{})
	database.MigrateModels(db, &models.User{})
}

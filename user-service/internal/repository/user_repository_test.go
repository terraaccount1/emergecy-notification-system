package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/terraaccount1/emergecy-notification-system/user-service/internal/models"
	"github.com/terraaccount1/emergecy-notification-system/user-service/pkg/database"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	repo *UserRepository
)

func init() {
	dsn := "host=localhost user=test password=test dbname=test port=5432 sslmode=disable"
	db = database.NewPostgresConnection(dsn)
	repo = NewUserRepository(db)
}

func TestCreateUserWithContacts(t *testing.T) {

	database.MigrateModels(db, &models.User{})
	database.MigrateModels(db, &models.Contact{})

	data := models.User{
		Firstname: "A",
		Surname:   "A",
		Contacts: []models.Contact{
			{
				Firstname: "A",
				Surname:   "A",
				Phone:     "+3731234567890",
				Email:     "example@gmail.com",
			},
			{
				Firstname: "B",
				Surname:   "B",
				Phone:     "+3731234567890",
				Email:     "example@gmail.com",
			},
			{
				Firstname: "C",
				Surname:   "C",
				Phone:     "+3731234567890",
				Email:     "example@gmail.com",
			},
		},
	}

	result := models.User{
		ID:        1,
		Firstname: "A",
		Surname:   "A",
		Contacts: []models.Contact{
			{
				ID:        1,
				Firstname: "A",
				Surname:   "A",
				Email:     "example@gmail.com",
				Phone:     "+3731234567890",
				UserID:    1,
			},
			{
				ID:        2,
				Firstname: "B",
				Surname:   "B",
				Email:     "example@gmail.com",
				Phone:     "+3731234567890",
				UserID:    1,
			},
			{
				ID:        3,
				Firstname: "C",
				Surname:   "C",
				Email:     "example@gmail.com",
				Phone:     "+3731234567890",
				UserID:    1,
			},
		},
	}
	err := repo.CreateUserWithContacts(&data)

	assert.NoError(t, err)
	assert.Equal(t, result, data)

	repo.db.Migrator().DropTable(&models.Contact{})
	repo.db.Migrator().DropTable(&models.User{})
}

func TestCreateUserWithEmptyContacts(t *testing.T) {

	database.MigrateModels(db, &models.User{})
	database.MigrateModels(db, &models.Contact{})

	data := models.User{
		Firstname: "A",
		Surname:   "A",
		Contacts:  []models.Contact{},
	}

	result := models.User{
		ID:        1,
		Firstname: "A",
		Surname:   "A",
		Contacts:  []models.Contact{},
	}
	err := repo.CreateUserWithContacts(&data)

	assert.NoError(t, err)
	assert.Equal(t, result, data)

	repo.db.Migrator().DropTable(&models.Contact{})
	repo.db.Migrator().DropTable(&models.User{})
}

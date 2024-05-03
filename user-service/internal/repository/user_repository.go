package repository

import (
	"github.com/terraaccount1/emergecy-notification-system/user-service/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) CreateUserWithContacts(user *models.User) error {
	return repo.db.Model(&models.User{}).Create(user).Error
}

func (repo *UserRepository) CreateUserContacts(contacts *[]models.Contact) error {
	return repo.db.Model(&models.Contact{}).Create(contacts).Error
}

func (repo *UserRepository) ReadUserWithContacts(user *models.User) error {
	if err := repo.db.Model(&models.User{}).Where(&models.User{ID: user.ID}).First(user).Error; err != nil {
		return err
	}
	return repo.db.Model(&models.Contact{}).Where(&models.Contact{UserID: user.ID}).Find(user.Contacts).Error
}

func (repo *UserRepository) UpdateUserContact(contact *models.Contact) error {
	return repo.db.Model(&models.Contact{}).Where(models.Contact{Firstname: contact.Firstname, Surname: contact.Surname}).Updates(contact).Error
}

func (repo *UserRepository) DeleteUserContact(contact *models.Contact) error {
	return repo.db.Model(&models.Contact{}).Where(&models.Contact{Firstname: contact.Firstname, Surname: contact.Surname}).Delete(contact).Error
}

func (repo *UserRepository) DeleteUserWithContacts(user *models.User) error {
	if err := repo.db.Model(&models.User{}).Where(&models.User{ID: user.ID}).Delete(user).Error; err != nil {
		return err
	}
	return repo.db.Model(&models.Contact{}).Where(&models.Contact{UserID: user.ID}).Delete(user.Contacts).Error
}

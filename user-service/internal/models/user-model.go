package models

type User struct {
	ID uint
	// gorm.Model
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`

	Contacts []Contact
}

type Contact struct {
	ID uint
	// gorm.Model
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`

	Email string `json:"email" example:"example@abc.com"`
	Phone string `json:"phone" example:"+373 60-000-000"`

	UserID uint `gorm:"index"`
}

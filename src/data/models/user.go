package models

type User struct {
	BaseModel
	ID             uint   `gorm:"primaryKey"`
	Username       string `gorm:"size:50;unique;not null"`
	Email          string `gorm:"size:100;unique;not null"`
	Password       string `gorm:"size:255;not null"` // Hashed password
	FullName       string `gorm:"size:100"`
	ProfilePicture string `gorm:"size:255"` // URL to profile picture
}

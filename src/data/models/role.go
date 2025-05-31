package models

type Role struct {
	BaseModel
	Name     string `gorm:"type:varchar(50);not null;unique"`
	UserRole *[]UserRole
}

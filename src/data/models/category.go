package models

type Category struct {
	BaseModel
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:50;unique;not null"`
	Description string `gorm:"type:text"`
	Posts       []Post `gorm:"foreignKey:CategoryID"`
}

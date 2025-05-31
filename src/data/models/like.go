package models

type Like struct {
	BaseModel
	ID     uint `gorm:"primaryKey"`
	PostID uint `gorm:"not null;index"`
	Post   Post `gorm:"foreignKey:PostID"`
	UserID uint `gorm:"not null;index"`
	User   User `gorm:"foreignKey:UserID"`
}

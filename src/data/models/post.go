package models

type Post struct {
	BaseModel
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null;index"`
	User       User      `gorm:"foreignKey:UserID"`
	Title      string    `gorm:"size:255;not null"`
	Content    string    `gorm:"type:text;not null"`
	CategoryID uint      `gorm:"not null;index"`
	Category   Category  `gorm:"foreignKey:CategoryID"`
	Status     string    `gorm:"size:20;default:'draft'"` // draft, published
	Comments   []Comment `gorm:"foreignKey:PostID"`
	Likes      []Like    `gorm:"foreignKey:PostID"`
}

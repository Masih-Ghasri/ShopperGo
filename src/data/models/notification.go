package models

type Notification struct {
	BaseModel
	ID            uint   `gorm:"primaryKey"`
	UserID        uint   `gorm:"not null;index"`
	User          User   `gorm:"foreignKey:UserID"`
	Type          string `gorm:"size:20;not null"` // like, comment, follow
	ReferenceID   uint   `gorm:"not null;index"`   // ID of Post/Comment/Follow
	ReferenceType string `gorm:"size:20;not null"` // post, comment, follow
	IsRead        bool   `gorm:"default:false"`
}

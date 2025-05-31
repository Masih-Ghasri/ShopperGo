package models

type Follow struct {
	BaseModel
	ID         uint `gorm:"primaryKey"`
	FollowerID uint `gorm:"not null;index"`
	Follower   User `gorm:"foreignKey:FollowerID"`
	FollowedID uint `gorm:"not null;index"`
	Followed   User `gorm:"foreignKey:FollowedID"`
}

package models

type User struct {
	BaseModel
	ID             uint   `gorm:"primaryKey"`
	Username       string `gorm:"size:20;unique;not null;type:varchar(20)"`
	Email          string `gorm:"size:100;unique;not null;type:varchar(100)"`
	Password       string `gorm:"size:255;not null;type:varchar(255)"` // Hashed password
	Name           string `gorm:"size:20;unique;not null;type:varchar(20)"`
	LastName       string `gorm:"size:20;unique;not null;type:varchar(20)"`
	PhoneNumber    string `gorm:"size:11;unique;default:null;type:varchar(11)"`
	ProfilePicture string `gorm:"size:255;type:varchar(255)"` // URL to profile picture
	Bio            string `gorm:"type:text"`
	UserRole       *[]UserRole
	Posts          []Post         `gorm:"foreignKey:UserID"`
	Comments       []Comment      `gorm:"foreignKey:UserID"`
	Likes          []Like         `gorm:"foreignKey:UserID"`
	Followers      []Follow       `gorm:"foreignKey:FollowedID"` // Users who follow this user
	Following      []Follow       `gorm:"foreignKey:FollowerID"` // Users this user follows
	Notifications  []Notification `gorm:"foreignKey:UserID"`
}

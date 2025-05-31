package models

type UserRole struct {
	BaseModel
	User   User `gorm:"foreignkey:UserID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	Role   Role `gorm:"foreignkey:RoleID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	UserID int
	RoleID int
}

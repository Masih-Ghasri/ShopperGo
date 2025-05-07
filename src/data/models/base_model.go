package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id         int            `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt  time.Time      `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP;not null"`
	ModifiedAt sql.NullTime   `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP;null"`
	DeletedAt  sql.NullTime   `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP;null"`
	CreatedBy  int            `gorm:"not null"`
	ModifiedBy *sql.NullInt64 `gorm:"null"`
	DeletedBy  *sql.NullInt64 `gorm:"null"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId = -1
	if value != nil {
		userId = int(value.(float64))
	}
	m.CreatedAt = time.Now().UTC()
	m.CreatedBy = userId
	return nil
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId = &sql.NullInt64{Valid: false}
	if value != nil {
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}
	m.ModifiedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	m.ModifiedBy = userId
	return nil
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId = &sql.NullInt64{Valid: false}
	if value != nil {
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}
	m.DeletedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	m.DeletedBy = userId
	return nil
}

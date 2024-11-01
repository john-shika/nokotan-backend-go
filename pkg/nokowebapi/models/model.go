package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint64         `db:"id" gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time      `db:"created_at" gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time      `db:"updated_at" gorm:"not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" gorm:"index" json:"deletedAt"`
}

func (Model) TableName() string {
	return "models"
}

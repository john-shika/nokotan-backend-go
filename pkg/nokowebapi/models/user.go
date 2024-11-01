package models

import "database/sql"

type User struct {
	Model
	UUID     string         `db:"uuid" gorm:"unique;not null;index" json:"uuid"`
	Username string         `db:"username" gorm:"unique;not null;index" json:"username"`
	Password string         `db:"password" gorm:"not null" json:"password"`
	Email    sql.NullString `db:"email" gorm:"unique;index" json:"email,omitempty"`
	Phone    sql.NullString `db:"phone" gorm:"unique;index" json:"phone,omitempty"`
	Admin    string         `db:"admin" gorm:"not null;index" json:"admin"`
	Role     string         `db:"role" gorm:"not null;index" json:"role"`
}

func (User) TableName() string {
	return "users"
}

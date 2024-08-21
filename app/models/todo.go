package models

type Todo struct {
	ID          int    `db:"id" gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string `db:"title" gorm:"unique;not null" json:"title"`
	Description string `db:"description" gorm:"nullable" json:"description"`
}

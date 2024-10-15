package models

type Todo struct {
	Model
	Title       string `db:"title" gorm:"unique;not null" json:"title,required"`
	Description string `db:"description" gorm:"nullable" json:"description,required"`
}

func (Todo) TableName() string {
	return "todos"
}

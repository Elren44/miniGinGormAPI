package models

import (
	"gorm.io/gorm"
)

type ToDo struct {
	gorm.Model
	ToDoTime string `json:"todo_time"`
	Content  string `json:"content"`
}

func (a *ToDo) TableName() string {
	return "todo_list"
}

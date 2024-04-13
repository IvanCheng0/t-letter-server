package model

import (
	"database/sql"
	"time"
)

type GormModel struct {
	ID        uint         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" gorm:"index"`
}

type TodoList struct {
	GormModel
	Content string `json:"content"`
	Completed bool `json:"completed"`
}

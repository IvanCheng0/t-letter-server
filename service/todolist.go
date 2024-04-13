package service

import (
	"github.com/IvanCheng0/t-letter-server/model"
	"gorm.io/gorm"
)

type TodoList struct {
	*gorm.DB
}

func (list *TodoList) GetList() []model.TodoList {
	var todos []model.TodoList
	list.Order("id desc").Find(&todos)
	return todos
}

func (list *TodoList) AddTodo(content string) uint {
	todo := model.TodoList{
		Content:   content,
		Completed: false,
	}
	list.Create(&todo)
	return todo.ID
}

func (list *TodoList) ToggleTodo(id uint) uint {
	var completed bool
	list.Model(&model.TodoList{}).Where("id = ?", id).Select("completed").Scan(&completed)
	list.Model(&model.TodoList{}).Where("id = ?", id).Update("completed", !completed)

	return id
}

func (list *TodoList) DeleteTodo(id uint) uint {
	list.Delete(&model.TodoList{}, id)

	return id
}

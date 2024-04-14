package controller

import (
	"net/http"

	"github.com/IvanCheng0/t-letter-server/service"
	"github.com/gin-gonic/gin"
)

type PostBody struct {
	Id      uint
	Content string
}

type TodoList struct {
	// *TodoList
	service.TodoList
}

func (list *TodoList) Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Methods", "GET, POST")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, "")
		c.Abort()
		return
	}
	c.Next()
}

func (list *TodoList) List(c *gin.Context) {
	result := list.GetList()
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "data": result})
}

func (list *TodoList) Add(c *gin.Context) {
	var body PostBody
	c.ShouldBindJSON(&body)
	result := list.AddTodo(body.Content)
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "data": result})

}

func (list *TodoList) Toggle(c *gin.Context) {
	var body PostBody
	c.ShouldBindJSON(&body)
	result := list.ToggleTodo(body.Id)
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "data": result})
}

func (list *TodoList) Delete(c *gin.Context) {
	var body PostBody
	c.ShouldBindJSON(&body)
	result := list.ToggleTodo(body.Id)
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "data": result})
}

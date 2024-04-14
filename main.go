package main

import (
	"fmt"

	"github.com/IvanCheng0/t-letter-server/controller"
	"github.com/IvanCheng0/t-letter-server/db"
	"github.com/IvanCheng0/t-letter-server/model"
	"github.com/IvanCheng0/t-letter-server/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("hello Gin")
	// g := gin.Default()

	// g.GET("/test", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{
	// 				"message": "Hello, Gin!",
	// 		})
	// })

	// g.Run(":8081")
	r := gin.Default()
	db := db.Connect("test")
	fmt.Println(db)
	// TODO: 这里需要看下
	db.AutoMigrate(&model.TodoList{})
	todoListService := service.TodoList{DB: db,}
	todoListController := controller.TodoList{TodoList: todoListService}
	
	r.Use(todoListController.Cors)
	r.GET("/list", todoListController.List)
    r.POST("/add", todoListController.Add)
	r.POST("/toggle", todoListController.Toggle)
	r.POST("/delete", todoListController.Delete)

	r.Run(":3000")

}

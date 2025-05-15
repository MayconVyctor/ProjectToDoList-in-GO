package main

import (
	"todo-api/controller"
	"todo-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	TaskUseCase := usecase.NewTaskUseCase()
	TaskControler := controller.NewTaskController(TaskUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/tasks", TaskControler.GetTasks)

	server.Run(":8000")
}

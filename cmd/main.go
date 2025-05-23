package main

import (
	"todo-api/controller"
	"todo-api/db"
	"todo-api/repository"
	"todo-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	TaskRepository := repository.NewTaskRepository(dbConnection)

	TaskUseCase := usecase.NewTaskUseCase(TaskRepository)

	TaskControler := controller.NewTaskController(TaskUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/tasks", TaskControler.GetTasks)
	server.POST("/task", TaskControler.CreateTask)

	server.Run(":8000")
}

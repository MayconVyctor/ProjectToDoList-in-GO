package controller

import (
	"net/http"
	"time"
	"todo-api/model"
	"todo-api/usecase"

	"github.com/gin-gonic/gin"
)

type taskController struct {
	taskUseCase usecase.TaskUsecase
}

func NewTaskController(usecase usecase.TaskUsecase) taskController {
	return taskController{
		taskUseCase: usecase,
	}
}

func (t *taskController) GetTasks(ctx *gin.Context) {

	tasks := []model.Task{
		{
			ID:          1,
			Title:       "Jogar volei",
			Description: "treino volei hoje as 19:30",
			Completed:   false,
			CreatedAt:   time.Now(),
		},
	}

	ctx.JSON(http.StatusOK, tasks)
}

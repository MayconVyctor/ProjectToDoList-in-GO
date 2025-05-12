package task

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var tasks = []Task{}

func CreateTask(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados invalidos"})
		return
	}

	newTask.Id = len(tasks) + 1
	newTask.CreatedAt = time.Now()
	tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, newTask)
}

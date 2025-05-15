package model

import (
	"time"
)

type Task struct {
	ID          int       `json:"id_task"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
}

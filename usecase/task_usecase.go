package usecase

import "todo-api/model"

type TaskUsecase struct {
}

func NewTaskUseCase() TaskUsecase {
	return TaskUsecase{}
}

func (tu *TaskUsecase) GetTask() ([]model.Task, error) {
	return []model.Task{}, nil
}

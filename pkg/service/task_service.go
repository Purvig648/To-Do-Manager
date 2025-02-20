package service

import (
	"github.com/todo_manager/pkg/model"
	dbmodel "github.com/todo_manager/pkg/model/db_model"
)

func (s *service) CreateTask(id uint, req model.Task) (dbmodel.Task, int, error) {
	taskData := dbmodel.Task{
		TaskName:        req.TaskName,
		TaskDescription: req.TaskDescription,
		TaskDeadline:    req.TaskDeadline,
		TaskStatus:      req.TaskStatus,
		UserID:          id,
	}
	createdTask, statusCode, err := s.repo.CreateTask(id, taskData)
	if err != nil {
		return dbmodel.Task{}, statusCode, err
	}
	return createdTask, statusCode, nil
}

package service

import (
	"github.com/rs/zerolog/log"
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

func (s *service) ViewAllTask() ([]model.TaskResponse, int, error) {
	allTaskDetails, statusCode, err := s.repo.ViewAllTasK()
	if err != nil {
		log.Error().Err(err)
		return nil, statusCode, err
	}
	taskDetail := []model.TaskResponse{}
	for i := 0; i < len(allTaskDetails); i++ {
		taskDetail = append(taskDetail, model.TaskResponse{
			ID:              allTaskDetails[i].ID,
			TaskName:        allTaskDetails[i].TaskName,
			TaskDescription: allTaskDetails[i].TaskDescription,
			TaskDeadline:    allTaskDetails[i].TaskDeadline,
			TaskStatus:      allTaskDetails[i].TaskStatus,
			UserID:          allTaskDetails[i].UserID,
		})
	}
	return taskDetail, statusCode, nil
}

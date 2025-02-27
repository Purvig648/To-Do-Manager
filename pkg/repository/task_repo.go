package repository

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/todo_manager/pkg/model"
	dbmodel "github.com/todo_manager/pkg/model/db_model"
	"gorm.io/gorm"
)

func (r *repo) CreateTask(uid uint, taskData dbmodel.Task) (dbmodel.Task, int, error) {
	taskData.UserID = uid
	tx := r.db.Create(&taskData)
	if err := tx.Error; err != nil {
		log.Error().Err(err)
		return dbmodel.Task{}, http.StatusBadRequest, err
	}
	return taskData, http.StatusAccepted, nil
}

func (r *repo) ViewAllTasK() ([]dbmodel.Task, int, error) {
	var allTasks []dbmodel.Task
	tx := r.db.Find(&allTasks)
	if err := tx.Error; err != nil {
		log.Error().Err(err)
		return nil, http.StatusBadRequest, err
	}
	return allTasks, http.StatusAccepted, nil
}

func (r *repo) ViewAllTaskOfUser(uid uint) ([]dbmodel.Task, int, error) {
	var alluserTasks []dbmodel.Task
	tx := r.db.Where("user_id", uid).Find(&alluserTasks)
	if err := tx.Error; err != nil {
		log.Error().Err(err)
		return nil, http.StatusBadRequest, err
	}
	return alluserTasks, http.StatusAccepted, nil
}

func (r *repo) ViewTask(tid uint) (dbmodel.Task, int, error) {
	var taskDetail dbmodel.Task
	tx := r.db.Where("id", tid).First(&taskDetail)
	if err := tx.Error; err != nil {
		log.Error().Err(err)
		return dbmodel.Task{}, http.StatusBadRequest, err
	}
	return taskDetail, http.StatusAccepted, nil
}

func (r *repo) UpdateTaskStatus(tid uint, choice string) (dbmodel.Task, int, error) {
	taskStatusDetail := dbmodel.Task{
		Model: gorm.Model{
			ID: tid,
		},
	}
	tx := r.db.Model(&taskStatusDetail).Updates(dbmodel.Task{
		TaskStatus: choice,
	})
	if err := tx.Error; err != nil {
		log.Error().Err(err)
		return dbmodel.Task{}, http.StatusBadRequest, err
	}
	return taskStatusDetail, http.StatusAccepted, nil
}

func (r *repo) UpadteAllTaskDetail(tid uint, taskUpdateDetails model.TaskDetailsUpdate) (dbmodel.Task, int, error) {
	taskUpdateDetail := dbmodel.Task{
		Model: gorm.Model{
			ID: tid,
		},
	}
	tx := r.db.Model(&taskUpdateDetail).Updates(dbmodel.Task{
		TaskName:        taskUpdateDetails.TaskName,
		TaskDescription: taskUpdateDetails.TaskDescription,
		TaskDeadline:    taskUpdateDetails.TaskDeadline,
		TaskStatus:      taskUpdateDetails.TaskStatus,
	})
	if err := tx.Error; err != nil {
		log.Error().Err(err)
		return dbmodel.Task{}, http.StatusBadRequest, err
	}
	return taskUpdateDetail, http.StatusAccepted, nil
}

func (r *repo) DeleteTask(tid uint) (int, error) {
	task := dbmodel.Task{
		Model: gorm.Model{
			ID: tid,
		},
	}
	tx := r.db.Delete(&task)
	if err := tx.Error; err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusAccepted, nil
}

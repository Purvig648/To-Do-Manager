package repository

import (
	"net/http"

	"github.com/rs/zerolog/log"
	dbmodel "github.com/todo_manager/pkg/model/db_model"
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

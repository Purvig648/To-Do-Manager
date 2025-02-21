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

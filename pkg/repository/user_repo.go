package repository

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/todo_manager/pkg/model"
	dbmodel "github.com/todo_manager/pkg/model/db_model"
	"gorm.io/gorm"
)

func (r *repo) SignUpRepo(userData dbmodel.User) (uint, int, error) {
	tx := r.db.Create(&userData)
	if err := tx.Error; err != nil {
		log.Error().Err(err)
		return 0, http.StatusInternalServerError, err
	}
	return userData.ID, http.StatusAccepted, nil
}

func (r *repo) CheckEmail(email string) (dbmodel.User, int, error) {
	var userDetails dbmodel.User
	tx := r.db.Where("email_id=?", email).First(&userDetails)
	if err := tx.Error; err != nil {
		log.Error().Err(err)
		return dbmodel.User{}, http.StatusBadRequest, err
	}
	return userDetails, http.StatusAccepted, nil
}

func (r *repo) ViewAllUsers() ([]dbmodel.User, int, error) {
	var allUserDetails []dbmodel.User
	tx := r.db.Find(&allUserDetails)
	if err := tx.Error; err != nil {
		log.Error().Err(err)
		return nil, http.StatusBadRequest, err
	}
	return allUserDetails, http.StatusAccepted, nil

}

func (r *repo) ViewUser(request model.UserRequest) (dbmodel.User, int, error) {
	var userDetail dbmodel.User
	tx := r.db.Where("id = ?", request.UserID).First(&userDetail)
	if err := tx.Error; err != nil {
		log.Error().Err(err)
		return dbmodel.User{}, http.StatusNotFound, err
	}
	return userDetail, http.StatusAccepted, nil
}

func (r *repo) UpdateAllDetails(id uint, request model.UserDetailsUpdate) (dbmodel.User, int, error) {
	userDetail := dbmodel.User{
		Model: gorm.Model{
			ID: id,
		},
	}
	tx := r.db.Model(&userDetail).Updates(dbmodel.User{
		Username: request.Username,
		EmailID:  request.EmailID,
	})
	if err := tx.Error; err != nil {
		log.Error().Err(err)
		return dbmodel.User{}, http.StatusBadRequest, err
	}
	return userDetail, http.StatusAccepted, nil
}

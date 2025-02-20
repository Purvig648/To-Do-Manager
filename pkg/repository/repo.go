package repository

import (
	"github.com/todo_manager/pkg/model"
	dbmodel "github.com/todo_manager/pkg/model/db_model"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

type RepoInterface interface {
	SignUpRepo(userData dbmodel.User) (dbmodel.User, int, error)
	CheckEmail(email string) (dbmodel.User, int, error)
	ViewAllUsers() ([]dbmodel.User, int, error)
	ViewUser(request model.UserRequest) (dbmodel.User, int, error)
	UpdateAllDetails(id uint, request model.UserDetailsUpdate) (dbmodel.User, int, error)
	UpdateDetailUsername(id uint, req model.UserDetailUpdate) (dbmodel.User, int, error)
	UpdateDetailEmail(id uint, req model.UserDetailUpdate) (dbmodel.User, int, error)
}

func NewRepoLayer(db *gorm.DB) RepoInterface {
	return &repo{
		db: db,
	}

}

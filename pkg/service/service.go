package service

import (
	"github.com/todo_manager/pkg/model"
	dbmodel "github.com/todo_manager/pkg/model/db_model"
	"github.com/todo_manager/pkg/repository"
)

type service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	SignUpService(signUpData model.SignUp) (dbmodel.User, int, error)
	SignInService(signInData model.SignIn) (int, error)
	ViewAllUsers() ([]model.UserResponse, int, error)
	ViewUser(request model.UserRequest) (model.UserResponse, int, error)
	UpdateAllDetails(id uint, req model.UserDetailsUpdate) (model.UserResponse, int, error)
	UpdateDetail(uid uint, req model.UserDetailUpdate, choice string) (model.UserResponse, int, error)

	CreateTask(id uint, req model.Task) (dbmodel.Task, int, error)
}

func NewServiceLayer(repo repository.RepoInterface) ServiceInterface {
	return &service{
		repo: repo,
	}
}

package service

import (
	"github.com/todo_manager/pkg/model"
	"github.com/todo_manager/pkg/repository"
)

type service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	SignUpService(signUpData model.SignUp) (uint, int, error)
	SignInService(signInData model.SignIn) (int, error)
	ViewAllUsers() ([]model.UserResponse, int, error)
	ViewUser(request model.UserRequest) (model.UserResponse, int, error)
	UpdateAllDetails(id uint, req model.UserDetailsUpdate) (model.UserResponse, int, error)
}

func NewServiceLayer(repo repository.RepoInterface) ServiceInterface {
	return &service{
		repo: repo,
	}
}

package service

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/rs/zerolog/log"
	"github.com/todo_manager/pkg/model"
	dbmodel "github.com/todo_manager/pkg/model/db_model"
	"github.com/todo_manager/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUpService(signUpData model.SignUp) (dbmodel.User, int, error) {
	if signUpData.EmailID == "" {
		log.Error().Err(errors.New("email_id is required")).Msg("email_id is empty")
		return dbmodel.User{}, http.StatusBadRequest, errors.New("emailID is empty")
	} else {
		re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !re.MatchString(signUpData.EmailID) {
			log.Error().Err(errors.New("email_id is not valid")).Msg("email_id is not valid")
			return dbmodel.User{}, http.StatusBadRequest, errors.New("Invalid emailID")
		}
	}
	if signUpData.Password != signUpData.ConfirmPassword {
		log.Error().Err(errors.New("the passwords don't match")).Msg("the passwords don't match")
		return dbmodel.User{}, http.StatusBadRequest, errors.New("different passwords")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpData.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(errors.New("cannot hash password")).Msg("cannot hash the password")
		return dbmodel.User{}, http.StatusBadRequest, err
	}
	dbUserData := dbmodel.User{
		Username:       signUpData.Username,
		EmailID:        signUpData.EmailID,
		HashedPassword: string(hashedPassword),
	}

	userData, statusCode, err := s.repo.SignUpRepo(dbUserData)
	if err != nil {
		return dbmodel.User{}, statusCode, err
	}
	return userData, statusCode, nil
}

func (s *service) SignInService(signInData model.SignIn) (int, error) {
	userDetails, statusCode, err := s.repo.CheckEmail(signInData.EmailID)
	if err != nil {
		log.Error().Err(err).Msg("email id is not correct")
		return statusCode, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userDetails.HashedPassword), []byte(signInData.Password))
	if err != nil {
		log.Error().Err(errors.New("the entered password is incorrect")).Msg("password does not match")
		return http.StatusBadRequest, err
	}
	return statusCode, nil
}

func (s *service) ViewAllUsers() ([]model.UserResponse, int, error) {
	allUserDetails, statuscode, err := s.repo.ViewAllUsers()
	if err != nil {
		log.Error().Err(err)
		return nil, statuscode, err
	}
	userDetails := []model.UserResponse{}
	for i := 0; i < len(allUserDetails); i++ {
		userDetails = append(userDetails, model.UserResponse{
			Username: allUserDetails[i].Username,
			EmailID:  allUserDetails[i].EmailID,
		})
	}
	return userDetails, statuscode, nil
}

func (s *service) ViewUser(request model.UserRequest) (model.UserResponse, int, error) {
	userDetails, statusCode, err := s.repo.ViewUser(request)
	if err != nil {
		log.Error().Err(err)
		return model.UserResponse{}, statusCode, err
	}
	return model.UserResponse{
		Username: userDetails.Username,
		EmailID:  userDetails.EmailID,
	}, statusCode, nil
}

func (s *service) UpdateAllDetails(id uint, req model.UserDetailsUpdate) (model.UserResponse, int, error) {
	userDetailsUpdated, statusCode, err := s.repo.UpdateAllDetails(id, req)
	if err != nil {
		log.Error().Err(err)
		return model.UserResponse{}, statusCode, err
	}
	return model.UserResponse{
		Username: userDetailsUpdated.Username,
		EmailID:  userDetailsUpdated.EmailID,
	}, statusCode, nil
}

func (s *service) UpdateDetail(uid uint, req model.UserDetailUpdate, choice string) (model.UserResponse, int, error) {
	val, _ := util.Choices[choice]
	if val == "username" {
		resp, ststausCode, err := s.repo.UpdateDetailUsername(uid, req)
		if err != nil {
			log.Error().Err(err)
			return model.UserResponse{}, ststausCode, err
		}
		return model.UserResponse{
			Username: resp.Username,
		}, ststausCode, nil
	} else if val == "email_id" {
		resp, statusCode, err := s.repo.UpdateDetailEmail(uid, req)
		if err != nil {
			log.Error().Err(err)
			return model.UserResponse{}, statusCode, err
		}
		return model.UserResponse{
			EmailID: resp.EmailID,
		}, statusCode, nil
	}
	return model.UserResponse{}, http.StatusAccepted, nil
}

package service

import (
	"user-service/errs"
	"user-service/repositories"
)

type NewUser struct {
	Username string `gorm:"unique"`
	Password string
	Email    string
	LineID   string
}

type UserResponse struct {
	Username string `gorm:"unique"`
	Email    string
	LineID   string
}

type UserService interface {
	CreateAccount(user NewUser) (*UserResponse, error)
	GetUser(username string) (*UserResponse, error)
	DeleteAccount(username string) error
	UpdateInformation()
	UpdatePassword()
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return userService{userRepo}
}

func (s userService) CreateAccount(user NewUser) (*UserResponse, error) {
	if user.Username == "" || user.Password == "" {
		return nil, errs.InvalidRequest
	}

	users, err := s.userRepo.GetUsers()
	if err != nil {
		return nil, errs.ServerError
	}
	for _, u := range users {
		if u.Username == user.Username {
			return nil, errs.DuplicateUsername
		}
	}
	newAccount := repositories.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		LineID:   user.LineID,
	}
	s.userRepo.Create(newAccount)
	userResonse := UserResponse{
		Username: user.Username,
		Email:    user.Email,
		LineID:   user.LineID,
	}
	return &userResonse, nil
}

func (s userService) GetUser(username string) (*UserResponse, error) {
	if username == "" {
		return nil, errs.InvalidRequest
	}

	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, errs.NotFound
	}

	userResponse := UserResponse{
		Username: user.Username,
		LineID:   user.LineID,
		Email:    user.Email,
	}
	return &userResponse, nil
}

func (s userService) DeleteAccount(username string) (err error) {
	err = s.userRepo.Delete(username)
	if err != nil {
		return errs.ServerError
	}
	return nil
}

package services

import (
	"user-service/errs"
	"user-service/repositories"

	"golang.org/x/crypto/bcrypt"
)

type NewUser struct {
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Email    string
	LineID   string
}

type UpdateUserInfo struct {
	Username string
	Email    string
	LineID   string
}

type UpdateUserPassword struct {
	Username string
	Password string
}

type LoginUser struct {
	Username string
	Password string
}

type UserResponse struct {
	Username string
	Email    string
	LineID   string
}

type UserService interface {
	CreateAccount(user NewUser) (*UserResponse, error)
	GetUser(username string) (*UserResponse, error)
	DeleteAccount(username string) error
	UpdateInformation(UpdateUserInfo) error
	UpdatePassword(UpdateUserPassword) error
	Login(LoginUser) error
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

	Password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, errs.ServerError
	}

	newAccount := repositories.User{
		Username: user.Username,
		Password: string(Password),
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
	if username == "" {
		return errs.InvalidRequest
	}
	err = s.userRepo.Delete(username)
	if err != nil {
		return errs.ServerError
	}
	return nil
}

func (s userService) UpdateInformation(userUpdate UpdateUserInfo) error {
	if userUpdate.Username == "" {
		return errs.InvalidRequest
	}
	userUpdateRepo := repositories.UpdateUserInformation{
		Username: userUpdate.Username,
		Email:    userUpdate.Email,
		LineID:   userUpdate.LineID,
	}
	err := s.userRepo.UpdateInformation(userUpdateRepo)
	if err != nil {
		return errs.ServerError
	}
	return nil
}

func (s userService) UpdatePassword(userUpdate UpdateUserPassword) error {
	if userUpdate.Username == "" {
		return errs.InvalidRequest
	}
	userUpdateRepo := repositories.UpdateUserPassword{
		Username: userUpdate.Username,
		Password: userUpdate.Password,
	}
	err := s.userRepo.UpdatePassword(userUpdateRepo)
	if err != nil {
		return errs.ServerError
	}
	return nil
}

func (s userService) Login(user LoginUser) error {
	if user.Username == "" || user.Password == "" {
		return errs.InvalidRequest
	}
	userData, err := s.userRepo.GetUserByUsername(user.Username)
	if err != nil {
		return errs.UserNotFound
	}
	if err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		return errs.PasswordInvalid
	}
	return nil
}

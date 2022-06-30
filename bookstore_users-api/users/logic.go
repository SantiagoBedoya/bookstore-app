package users

import (
	"net/http"
	"time"

	"github.com/SantiagoBedoya/bookstore_utils/resterrors"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) CreateUser(user *User) (*User, *resterrors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Status = "active"
	user.Password = string(hash)
	user.CreatedAt = time.Now()
	return s.repository.CreateUser(user)
}

func (s *service) GetUser(userId string) (*User, *resterrors.RestError) {
	return s.repository.GetUser(userId)
}

func (s *service) UpdateUser(userId string, user *User) (*User, *resterrors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	_, err := s.repository.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return s.repository.UpdateUser(userId, user)
}

func (s *service) DeleteUser(userId string) (*User, *resterrors.RestError) {
	_, err := s.repository.GetUser(userId)
	if err != nil {
		return nil, err
	}

	return s.repository.DeleteUser(userId)
}

func (s *service) GetUsersByStatus(status string) (Users, *resterrors.RestError) {
	return s.repository.GetUsersByStatus(status)
}

func (s *service) GetUsers() (Users, *resterrors.RestError) {
	return s.repository.GetUsers()
}

func (s *service) SignIn(user *User) (*User, *resterrors.RestError) {
	currentUser, err := s.repository.SearchUserByEmail(user.Email)
	if err != nil {
		if err.StatusCode == http.StatusNotFound {
			return nil, resterrors.NewBadRequestError(InvalidCredentialsErr.Error())
		}
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(user.Password)); err != nil {
		return nil, resterrors.NewBadRequestError(InvalidCredentialsErr.Error())
	}
	return currentUser, nil
}

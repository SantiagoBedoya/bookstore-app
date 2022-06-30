package accesstoken

import (
	"time"

	"github.com/SantiagoBedoya/bookstore_utils/resterrors"
)

const EXPIRATION_TIME = time.Hour * 24

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAccessTokenById(id string) (*AccessToken, *resterrors.RestError) {
	return s.repository.GetAccessTokenById(id)
}

func (s *service) CreateAccessToken(at *AccessToken) (*AccessToken, *resterrors.RestError) {
	at.Expires = time.Now().UTC().Add(EXPIRATION_TIME).Unix()
	return s.repository.CreateAccessToken(at)
}

func (s *service) UpdateAccessTokenExpirationTime(at *AccessToken) *resterrors.RestError {
	return s.repository.UpdateAccessTokenExpirationTime(at)
}

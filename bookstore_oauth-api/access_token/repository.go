package accesstoken

import "github.com/SantiagoBedoya/bookstore_utils/resterrors"

type Repository interface {
	GetAccessTokenById(id string) (*AccessToken, *resterrors.RestError)
	CreateAccessToken(at *AccessToken) (*AccessToken, *resterrors.RestError)
	UpdateAccessTokenExpirationTime(at *AccessToken) *resterrors.RestError
}

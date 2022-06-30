package accesstoken

import "errors"

var (
	GetAccessTokenErr      = errors.New("error when trying to get access token")
	CreateAccessTokenErr   = errors.New("error when trying to create access token")
	UpdateAccessTokenErr   = errors.New("error when trying to update access token")
	AccessTokenNotFoundErr = errors.New("access token not found")
)

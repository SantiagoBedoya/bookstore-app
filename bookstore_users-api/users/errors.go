package users

import "errors"

var (
	InvalidUserDataErr    = errors.New("invalid user data")
	InvalidCredentialsErr = errors.New("invalid credentials")
	UserNotFoundErr       = errors.New("user not found")
	UserAlreadyExistErr   = errors.New("user already exist with this email")
	SaveUserErr           = errors.New("error when trying to save user")
	UpdateUserErr         = errors.New("error when trying to update user")
	DeleteUserErr         = errors.New("error when trying to delete user")
	GetUsersErr           = errors.New("error when trying to get users")
	GetUserErr            = errors.New("error when tryinh to get user")
)

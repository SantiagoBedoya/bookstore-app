package users

import (
	"github.com/SantiagoBedoya/bookstore_utils/resterrors"
)

type Service interface {
	CreateUser(user *User) (*User, *resterrors.RestError)
	GetUser(userId string) (*User, *resterrors.RestError)
	GetUsers() (Users, *resterrors.RestError)
	GetUsersByStatus(status string) (Users, *resterrors.RestError)
	UpdateUser(userId string, user *User) (*User, *resterrors.RestError)
	DeleteUser(userId string) (*User, *resterrors.RestError)
	SignIn(user *User) (*User, *resterrors.RestError)
}

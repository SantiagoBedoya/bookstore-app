package users

import "github.com/SantiagoBedoya/bookstore_utils/resterrors"

type Repository interface {
	CreateUser(user *User) (*User, *resterrors.RestError)
	GetUser(userId string) (*User, *resterrors.RestError)
	GetUsers() ([]User, *resterrors.RestError)
	GetUsersByStatus(status string) ([]User, *resterrors.RestError)
	UpdateUser(userId string, user *User) (*User, *resterrors.RestError)
	DeleteUser(userId string) (*User, *resterrors.RestError)
	SearchUserByEmail(email string) (*User, *resterrors.RestError)
}

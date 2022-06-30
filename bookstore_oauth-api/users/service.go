package users

import "github.com/SantiagoBedoya/bookstore_utils/resterrors"

type Service interface {
	SignIn(user *SignInRequest) (*User, *resterrors.RestError)
}

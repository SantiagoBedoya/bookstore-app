package users

import (
	"time"

	"github.com/SantiagoBedoya/bookstore_utils/resterrors"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id        int64     `json:"id"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Status    string    `json:"status"`
	Password  string    `json:"password" validate:"required,min=8"`
	CreatedAt time.Time `json:"created_at"`
}
type Users []User

func (u *User) Validate() *resterrors.RestError {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return resterrors.NewBadRequestError(err.Error())
		}
	}
	return nil
}

func (u *User) ToPublicUser() PublicUser {
	return PublicUser{
		Id:        u.Id,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}

type PublicUser struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

func (u *Users) ToPublicUsers() []PublicUser {
	result := make([]PublicUser, 0)
	for _, user := range *u {
		result = append(result, PublicUser{
			Id:        user.Id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		})
	}
	return result
}

package accesstoken

import (
	"time"

	"github.com/SantiagoBedoya/bookstore_utils/resterrors"
	"github.com/go-playground/validator/v10"
)

type AccessToken struct {
	AccessToken string `json:"access_token" validate:"required"`
	UserId      int64  `json:"user_id" validate:"required"`
	ClientId    int64  `json:"client_id" validate:"required"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *resterrors.RestError {
	validate := validator.New()
	err := validate.Struct(at)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return resterrors.NewBadRequestError(err.Error())
		}
	}
	return nil
}

func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

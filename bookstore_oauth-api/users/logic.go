package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SantiagoBedoya/bookstore_utils/logger"
	"github.com/SantiagoBedoya/bookstore_utils/resterrors"
	"github.com/go-resty/resty/v2"
)

type service struct {
	baseURL string
	client  *resty.Client
}

func NewService(baseURL string) Service {
	return &service{
		baseURL: baseURL,
		client:  resty.New(),
	}
}

func (s *service) SignIn(user *SignInRequest) (*User, *resterrors.RestError) {
	url := fmt.Sprintf("%s/sign-in", s.baseURL)
	resp, err := s.client.R().SetBody(user).Post(url)
	if err != nil {
		logger.Error(err.Error())
		return nil, resterrors.NewRestError(BadGatewayUserErr.Error(), http.StatusBadGateway, "bad_gateway")
	}
	if resp.StatusCode() != http.StatusOK {
		var data resterrors.RestError
		if err := json.Unmarshal(resp.Body(), &data); err != nil {
			logger.Error(err.Error())
			return nil, resterrors.NewInternalServerError("invalid response")
		}
		return nil, &data
	}
	var currentUser User
	if err := json.Unmarshal(resp.Body(), &currentUser); err != nil {
		return nil, resterrors.NewInternalServerError("invalid response")
	}
	return &currentUser, nil
}

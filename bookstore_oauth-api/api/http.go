package api

import (
	"net/http"

	accesstoken "github.com/SantiagoBedoya/bookstore_oauth-api/access_token"
	"github.com/SantiagoBedoya/bookstore_oauth-api/users"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	SignIn(c *gin.Context)
}

type handler struct {
	service     accesstoken.Service
	userService users.Service
}

func NewHandler(service accesstoken.Service, userService users.Service) Handler {
	return &handler{service, userService}
}

func (h *handler) SignIn(c *gin.Context) {
	var data users.SignInRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	user, err := h.userService.SignIn(&data)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *handler) GetById(c *gin.Context) {
	atId := c.Param("id")
	at, err := h.service.GetAccessTokenById(atId)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, at)
}

func (h *handler) Create(c *gin.Context) {
	var at accesstoken.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	newAt, err := h.service.CreateAccessToken(&at)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusCreated, newAt)
}

func (h *handler) Update(c *gin.Context) {

}

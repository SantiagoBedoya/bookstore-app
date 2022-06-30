package api

import (
	"net/http"

	"github.com/SantiagoBedoya/bookstore_users-api/users"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUsersByStatus(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	SignIn(c *gin.Context)
}

type handler struct {
	service users.Service
}

func NewHandler(service users.Service) Handler {
	return &handler{service}
}

func (h *handler) SignIn(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	currentUser, err := h.service.SignIn(&user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, currentUser.ToPublicUser())
}

func (h *handler) GetUsersByStatus(c *gin.Context) {
	status := c.Query("status")
	is_public := c.GetBool("is_public")
	if status != "" {
		users, err := h.service.GetUsersByStatus(status)
		if err != nil {
			c.JSON(err.StatusCode, err)
			return
		}
		if is_public {
			data := users.ToPublicUsers()
			c.JSON(http.StatusOK, data)
			return
		}
		c.JSON(http.StatusOK, users)
		return
	}
	users, err := h.service.GetUsers()
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	if is_public {
		data := users.ToPublicUsers()
		c.JSON(http.StatusOK, data)
		return
	}
	c.JSON(http.StatusOK, users)

}

func (h *handler) CreateUser(c *gin.Context) {
	var user users.User
	is_public := c.GetBool("is_public")
	if is_public {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	newUser, err := h.service.CreateUser(&user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func (h *handler) GetUser(c *gin.Context) {
	userId := c.Param("user_id")
	user, err := h.service.GetUser(userId)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	is_public := c.GetBool("is_public")
	if is_public {
		data := user.ToPublicUser()
		c.JSON(http.StatusOK, data)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *handler) UpdateUser(c *gin.Context) {
	var user users.User
	userId := c.Param("user_id")
	is_public := c.GetBool("is_public")
	if is_public {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	_, err := h.service.UpdateUser(userId, &user)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *handler) DeleteUser(c *gin.Context) {
	is_public := c.GetBool("is_public")
	if is_public {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}
	userId := c.Param("user_id")
	_, err := h.service.DeleteUser(userId)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.Status(http.StatusNoContent)
}

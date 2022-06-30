package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/SantiagoBedoya/bookstore_users-api/api"
	"github.com/SantiagoBedoya/bookstore_users-api/repositories/mysql"
	"github.com/SantiagoBedoya/bookstore_users-api/users"
	"github.com/SantiagoBedoya/bookstore_utils/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	repository := mysql.NewMysqlRepository(os.Getenv("DB_CONNECTION"))
	router := gin.Default()

	r := router.Group("/api/users")
	usersRouter := r.Use(api.AuthMiddleware)
	{
		service := users.NewService(repository)
		handler := api.NewHandler(service)
		usersRouter.POST("", handler.CreateUser)
		usersRouter.GET("", handler.GetUsersByStatus)
		usersRouter.GET("/:user_id", handler.GetUser)
		usersRouter.PUT("/:user_id", handler.UpdateUser)
		usersRouter.DELETE("/:user_id", handler.DeleteUser)
		usersRouter.POST("/sign-in", handler.SignIn)
	}

	errs := make(chan error, 2)
	go func() {
		logger.Info(fmt.Sprintf("http://localhost%s", handlePort()))
		errs <- http.ListenAndServe(handlePort(), router)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("terminated %s\n", <-errs)
}

func handlePort() string {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	return fmt.Sprintf(":%s", port)
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	accesstoken "github.com/SantiagoBedoya/bookstore_oauth-api/access_token"
	"github.com/SantiagoBedoya/bookstore_oauth-api/api"
	"github.com/SantiagoBedoya/bookstore_oauth-api/repositories/cassandra"
	"github.com/SantiagoBedoya/bookstore_oauth-api/users"
	"github.com/SantiagoBedoya/bookstore_utils/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	repository, session := cassandra.NewCassandraRepository(os.Getenv("CASSANDRA_HOST"))
	defer session.Close()
	router := gin.Default()

	service := accesstoken.NewService(repository)
	userService := users.NewService(os.Getenv("USER_MS_URL"))
	handler := api.NewHandler(service, userService)

	r := router.Group("/oauth/access-token")
	{
		r.GET("/:id", handler.GetById)
		r.POST("", handler.Create)
		r.POST("/sign-in", handler.SignIn)
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

	fmt.Printf("terminated: %s\n", <-errs)
}

func handlePort() string {
	port := "4000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

package api

import "github.com/gin-gonic/gin"

func AuthMiddleware(c *gin.Context) {
	authorization := c.GetHeader("authorization")
	c.Set("is_public", authorization == "")
	c.Next()
}

package middleware

import "github.com/gin-gonic/gin"

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.Query("token")
		// Temporarily use user ID 1.
		c.Set("uid", "1")
		c.Next()
	}
}

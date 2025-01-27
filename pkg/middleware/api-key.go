package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.GetQuery("api")

		if !err {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
		}

		c.Next()
	}

}

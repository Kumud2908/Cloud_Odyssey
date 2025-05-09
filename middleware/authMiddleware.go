package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	helper "golang-hospital-management/helpers"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No authorization header provided")})
			c.Abort()
			return

		}
		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
		}
		c.Set("userClaims", claims)
		c.Next()
	}

}

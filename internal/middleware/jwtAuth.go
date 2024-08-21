package middleware

import (
	"github.com/emre-unlu/GinTest/internal/utils"
	"github.com/emre-unlu/GinTest/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Token ")
		if tokenString == authHeader {
			utils.RespondWithError(c, http.StatusUnauthorized, "Unable to find token")
			c.Abort()
			return
		}

		claims, err := jwt.ValidateJWT(tokenString)
		if err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("id", claims.Id)
		c.Next()
	}
}

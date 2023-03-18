package middleware

import (
	"github.com/gin-gonic/gin"
	"music_web/common"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "permission denied"})
			context.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "permission denied"})
			context.Abort()
			return
		}
		userId := claims.UserId
		context.Set("user", userId)
		context.Next()
	}
}

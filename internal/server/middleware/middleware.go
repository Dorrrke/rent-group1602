package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization") // Bearer <access_token>

		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		_ = strings.Split(authHeader, " ")[1]

		// uid, err := auth.ParseToken(tokenString)
		// if err != nil {
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		// 	return
		// }

		ctx.Set("uid", "uid")
		ctx.Next()
	}
}

package middleware

import (
	"net/http"
	"strings"

	"github.com/Dorrrke/rent-group1602/internal/service/auth"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization") // Bearer <access_token>

		log.Debug().Str("auth header", authHeader).Msg("auth middleware")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		uid, err := auth.ParseToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("uid", uid)
		ctx.Next()
	}
}

package middleware

import (
	"net/http"
	"strings"

	"github.com/Dorrrke/rent-group1602/internal/domain/users"
	"github.com/Dorrrke/rent-group1602/internal/service/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization") // Bearer <access_token>

		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		uid, role, err := auth.ParseToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("uid", uid)
		ctx.Set("role", role)
		ctx.Next()
	}
}

func RoleMiddleware(allowedRoles ...users.Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exists := ctx.Get("role")
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "no role"})
			return
		}

		userRole := role.(users.Role)
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				ctx.Next()
				return
			}
		}
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
	}
}

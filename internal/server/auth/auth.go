package auth

import (
	"net/http"

	"github.com/Dorrrke/rent-group1602/internal/service/auth"

	"github.com/gin-gonic/gin"
)

func Refresh(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "missing refresh token",
		})
		return
	}

	uid, role, err := auth.ParseToken(refreshToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAccessToken, err := auth.GenerateAccessToken(uid, role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}

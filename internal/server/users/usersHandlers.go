package users

import (
	"net/http"

	usersDomain "github.com/Dorrrke/rent-group1602/internal/domain/users"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	RegisterUser(req usersDomain.RegisterRequest) (string, error)
	LoginUser(req usersDomain.LoginRequest) (usersDomain.User, error)
}

type UsersHandler struct {
	userService UserService
}

func New(userService UserService) *UsersHandler {
	return &UsersHandler{
		userService: userService,
	}
}

func (h *UsersHandler) Login(ctx *gin.Context) {
	var req usersDomain.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.userService.LoginUser(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": "accessToken",
	})
}

func (h *UsersHandler) Register(ctx *gin.Context) {
	var req usersDomain.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := h.userService.RegisterUser(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"uid": uid})
}

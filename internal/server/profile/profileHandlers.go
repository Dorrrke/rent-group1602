package profile

import (
	"net/http"

	"github.com/Dorrrke/rent-group1602/internal/domain/cars"
	"github.com/Dorrrke/rent-group1602/internal/domain/users"
	"github.com/gin-gonic/gin"
)

type ProfileService interface {
	GetProfile(uid string) (users.User, error)
	GetHistory(uid string) ([]cars.Rent, error)
}

type ProfileHandlers struct {
	service ProfileService
}

func New(service ProfileService) *ProfileHandlers {
	return &ProfileHandlers{
		service: service,
	}
}

func (ph *ProfileHandlers) GetProfile(ctx *gin.Context) {
	userID, ok := ctx.Get("uid")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "missing token"})
		return
	}

	uid, ok := userID.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "missing uid"})
		return
	}

	user, err := ph.service.GetProfile(uid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (ph *ProfileHandlers) GetHistory(ctx *gin.Context) {
	userID, ok := ctx.Get("uid")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "missing token"})
		return
	}

	uid, ok := userID.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "missing uid"})
		return
	}

	history, err := ph.service.GetHistory(uid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, history)
}

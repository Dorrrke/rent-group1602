package cars

import (
	"net/http"

	carsDomain "github.com/Dorrrke/rent-group1602/internal/domain/cars"

	"github.com/gin-gonic/gin"
)

type CarsService interface {
	AddCar(req carsDomain.AddCarRequest) error
	GetAllCars() ([]carsDomain.Car, error)
	StartRent(uid string, req carsDomain.StartRentRequest) (float64, error)
	EndRent(req carsDomain.EndRentRequest) (float64, error)
}

type CarsHandlers struct {
	carsService CarsService
}

func New(carsService CarsService) *CarsHandlers {
	return &CarsHandlers{
		carsService: carsService,
	}
}

func (ch *CarsHandlers) AddCarHandler(ctx *gin.Context) {
	var req carsDomain.AddCarRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ch.carsService.AddCar(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Car added successfully"})
}

func (ch *CarsHandlers) GetAllCarsHandler(ctx *gin.Context) {
	cars, err := ch.carsService.GetAllCars()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, cars)
}

func (ch *CarsHandlers) StartRentHandler(ctx *gin.Context) {
	userID, ok := ctx.Get("uid")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "missing token"})
		return
	}

	uid := userID.(string)
	var req carsDomain.StartRentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	price, error := ch.carsService.StartRent(uid, req)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"price": price})
}

func (ch *CarsHandlers) EndRentHandler(ctx *gin.Context) {
	var req carsDomain.EndRentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	price, error := ch.carsService.EndRent(req)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"price": price})
}

package cars

import (
	carsDomain "github.com/Dorrrke/rent-group1602/internal/domain/cars"
)

// type CarsService interface {
// 	AddCar(req carsDomain.AddCarRequest) error
// 	GetAllCars() ([]carsDomain.Car, error)
// 	StartRent(uid string, req carsDomain.StartRentRequest) (float32, error)
// 	EndRent(uid string, req carsDomain.EndRentRequest) (float32, error)
// }

type Repository interface {
	AddCar(carsDomain.Car) error
	GetAllCars() ([]carsDomain.Car, error)
	StartRent(string, carsDomain.Rent) error
	EndRent(string, carsDomain.Rent) error
}

type CarService struct {
	repo Repository
}

func New(repo Repository) *CarService {
	return &CarService{
		repo: repo,
	}
}

func (s *CarService) AddCar(req carsDomain.AddCarRequest) error {
	panic("unimplemented")
}

func (s *CarService) GetAllCars() ([]carsDomain.Car, error) {
	panic("unimplemented")
}

func (s *CarService) StartRent(uid string, req carsDomain.StartRentRequest) (float32, error) {
	panic("unimplemented")
}

func (s *CarService) EndRent(uid string, req carsDomain.EndRentRequest) (float32, error) {
	panic("unimplemented")
}

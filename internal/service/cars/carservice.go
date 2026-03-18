package cars

import (
	carsDomain "github.com/Dorrrke/rent-group1602/internal/domain/cars"
	"github.com/Dorrrke/rent-group1602/internal/service/errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	AddCar(carsDomain.Car) error
	GetAllCars() ([]carsDomain.Car, error)
	StartRent(carsDomain.Rent) error
	EndRent(string) error

	GetCarByCID(cid string) (carsDomain.Car, error)
	GetRentByRID(rid string) (carsDomain.Rent, error)
}

type CarService struct {
	repo  Repository
	valid *validator.Validate
}

func New(repo Repository) *CarService {
	return &CarService{
		valid: validator.New(),
		repo:  repo,
	}
}

func (s *CarService) AddCar(req carsDomain.AddCarRequest) error {
	if err := s.valid.Struct(req); err != nil {
		return err
	}

	car := carsDomain.Car{
		CID:       uuid.NewString(),
		Brand:     req.Brand,
		Model:     req.Model,
		Color:     req.Color,
		Year:      req.Year,
		Number:    req.Number,
		Price:     req.Price,
		Available: true,
	}

	if err := s.repo.AddCar(car); err != nil {
		return err
	}

	return nil
}

func (s *CarService) GetAllCars() ([]carsDomain.Car, error) {
	return s.repo.GetAllCars()
}

func (s *CarService) StartRent(uid string, req carsDomain.StartRentRequest) (float64, error) {
	if err := s.valid.Struct(req); err != nil {
		return 0, err
	}

	rent := carsDomain.Rent{
		RID:   uuid.NewString(),
		UID:   uid,
		CID:   req.CID,
		Hours: req.Hours,
		Ended: false,
	}

	car, err := s.repo.GetCarByCID(rent.CID)
	if err != nil {
		return 0, err
	}

	if !car.Available {
		return 0, errors.ErrCarNotAvailable
	}

	rentPrice := car.Price * float64(rent.Hours)

	if err := s.repo.StartRent(rent); err != nil {
		return 0, err
	}

	return rentPrice, nil
}

func (s *CarService) EndRent(req carsDomain.EndRentRequest) (float64, error) {
	if err := s.valid.Struct(req); err != nil {
		return 0, err
	}

	rent, err := s.repo.GetRentByRID(req.RID)
	if err != nil {
		return 0, err
	}

	car, err := s.repo.GetCarByCID(rent.CID)
	if err != nil {
		return 0, err
	}

	totalPrice := float64(rent.Hours) * car.Price

	s.repo.EndRent(rent.RID)

	return totalPrice, nil
}

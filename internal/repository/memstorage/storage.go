package memstorage

import (
	carsDomain "github.com/Dorrrke/rent-group1602/internal/domain/cars"
	usersDomain "github.com/Dorrrke/rent-group1602/internal/domain/users"
	"github.com/Dorrrke/rent-group1602/internal/repository/errors"
)

type Storage struct {
	users map[string]usersDomain.User
	cars  map[string]carsDomain.Car
	rents map[string]carsDomain.Rent
}

func New() *Storage {
	return &Storage{
		users: make(map[string]usersDomain.User),
		cars:  make(map[string]carsDomain.Car),
		rents: make(map[string]carsDomain.Rent),
	}
}

func (s *Storage) SaveUser(user usersDomain.User) error {
	for _, u := range s.users {
		if u.Email == user.Email {
			return errors.ErrUserAlreadyExists
		}
	}

	s.users[user.UID] = user
	return nil
}

func (s *Storage) GetUserByEmail(email string) (usersDomain.User, error) {
	for _, u := range s.users {
		if u.Email == email {
			return u, nil
		}
	}

	return usersDomain.User{}, errors.ErrUserNotFound
}

func (s *Storage) GetUserByUID(uid string) (usersDomain.User, error) {
	user, ok := s.users[uid]

	if !ok {
		return usersDomain.User{}, errors.ErrUserNotFound
	}

	return user, nil
}

func (s *Storage) AddCar(car carsDomain.Car) error {
	for _, c := range s.cars {
		if c.Number == car.Number {
			return errors.ErrCarAlreadyExists
		}
	}

	s.cars[car.CID] = car
	return nil
}

func (s *Storage) GetAllCars() ([]carsDomain.Car, error) {
	cars := []carsDomain.Car{}
	for _, c := range s.cars {
		if c.Available {
			cars = append(cars, c)
		}
	}

	if len(cars) == 0 {
		return nil, errors.ErrNotAvailableCars
	}

	return cars, nil
}

func (s *Storage) GetCarByCID(cid string) (carsDomain.Car, error) {
	car, ok := s.cars[cid]

	if !ok {
		return carsDomain.Car{}, errors.ErrCarNotFound
	}

	return car, nil
}

func (s *Storage) GetRentByRID(rid string) (carsDomain.Rent, error) {
	rent, ok := s.rents[rid]

	if !ok {
		return carsDomain.Rent{}, errors.ErrRentNotFound
	}

	return rent, nil
}

func (s *Storage) StartRent(rent carsDomain.Rent) error {
	_, ok := s.rents[rent.RID]
	if ok {
		return errors.ErrRentAlreadyExists
	}

	s.rents[rent.RID] = rent
	return nil
}

func (s *Storage) EndRent(rid string) error {
	rent, ok := s.rents[rid]
	if !ok {
		return errors.ErrRentNotFound
	}

	rent.Ended = true
	s.rents[rid] = rent
	return nil
}

func (s *Storage) GetRentHistoryByID(uid string) ([]carsDomain.Rent, error) {
	history := []carsDomain.Rent{}

	for _, r := range s.rents {
		if r.UID == uid {
			history = append(history, r)
		}
	}

	if len(history) == 0 {
		return nil, errors.ErrEmptyHistory
	}

	return history, nil
}

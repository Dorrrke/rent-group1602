package memstorage

import (
	carsDomain "github.com/Dorrrke/rent-group1602/internal/domain/cars"
	usersDomain "github.com/Dorrrke/rent-group1602/internal/domain/users"
)

type Storage struct {
	users map[string]usersDomain.User
	cars  map[string]carsDomain.Car
	rents map[string]carsDomain.Rent
}

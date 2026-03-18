package profile

import (
	"github.com/Dorrrke/rent-group1602/internal/domain/cars"
	"github.com/Dorrrke/rent-group1602/internal/domain/users"
)

type ProfileRepository interface {
	GetUserByUID(string) (users.User, error)
	GetRentHistoryByID(string) ([]cars.Rent, error)
}

type ProfileService struct {
	repo ProfileRepository
}

func New(repo ProfileRepository) *ProfileService {
	return &ProfileService{
		repo: repo,
	}
}

func (u *ProfileService) GetProfile(uid string) (users.User, error) {
	return u.repo.GetUserByUID(uid)
}

func (u *ProfileService) GetHistory(uid string) ([]cars.Rent, error) {
	return u.repo.GetRentHistoryByID(uid)
}

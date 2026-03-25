package users

import (
	"fmt"

	usersDomain "github.com/Dorrrke/rent-group1602/internal/domain/users"
	"github.com/Dorrrke/rent-group1602/internal/service/errors"
	"github.com/google/uuid"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	SaveUser(user usersDomain.User) (uuid.UUID, error)
	GetUserByEmail(email string) (usersDomain.User, error)
	GetUserByUID(uid string) (usersDomain.User, error)
}

type UserService struct {
	repo  Repository
	valid *validator.Validate
}

func New(repo Repository) *UserService {
	return &UserService{
		repo:  repo,
		valid: validator.New(),
	}
}

func (s *UserService) RegisterUser(req usersDomain.RegisterRequest) (string, error) {
	if err := s.valid.Struct(req); err != nil {
		return "", fmt.Errorf(errors.IncorrectFieldValuesError, err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := usersDomain.User{
		Name:     req.Name,
		Age:      req.Age,
		Email:    req.Email,
		Password: string(hash),
		Role:     req.Role,
	}

	uid, err := s.repo.SaveUser(user)
	if err != nil {
		return "", err
	}

	return uid.String(), nil
}

func (s *UserService) LoginUser(req usersDomain.LoginRequest) (usersDomain.User, error) {
	if err := s.valid.Struct(req); err != nil {
		return usersDomain.User{}, fmt.Errorf(errors.IncorrectFieldValuesError, err)
	}

	storageUser, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return usersDomain.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storageUser.Password), []byte(req.Password))
	if err != nil {
		return usersDomain.User{}, errors.ErrInvalidCredentials
	}

	return storageUser, nil
}

package db

import (
	"context"
	"time"

	usersDomain "github.com/Dorrrke/rent-group1602/internal/domain/users"
	"github.com/google/uuid"
)

func (s *Storage) SaveUser(user usersDomain.User) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var uid uuid.UUID
	err := s.conn.QueryRow(
		ctx,
		`INSERT INTO users (name, age, email, password_hash, role, balance) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING uid`,
		user.Name, user.Age, user.Email,
		user.Password, user.Role, user.Balance,
	).Scan(&uid)

	if err != nil {
		// TODO: уточнение ошибки ( пользователь уже существует )
		return uuid.Nil, err
	}

	return uid, nil
}

func (s *Storage) GetUserByEmail(email string) (usersDomain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := s.conn.QueryRow(ctx, "SELECT uid, name, age, email, password_hash, role, balance FROM users WHERE email=$1", email)

	var user usersDomain.User
	if err := row.Scan(&user.UID, &user.Name, &user.Age, &user.Email, &user.Password, &user.Role, &user.Balance); err != nil {
		// TODO: уточнение ошибки ( нет такого пользователя )
		return usersDomain.User{}, err
	}

	return user, nil
}

func (s *Storage) GetUserByUID(uid string) (usersDomain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := s.conn.QueryRow(ctx, "SELECT * FROM users WHERE uid=$1", uid)

	var user usersDomain.User
	if err := row.Scan(&user.UID, &user.Name, &user.Age, &user.Email, &user.Password, &user.Role, &user.Balance); err != nil {
		// TODO: уточнение ошибки ( нет такого пользователя )
		return usersDomain.User{}, err
	}

	return usersDomain.User{}, nil
}

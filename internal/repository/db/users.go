package db

import (
	"context"
	"time"

	usersDomain "github.com/Dorrrke/rent-group1602/internal/domain/users"
)

func (s *Storage) SaveUser(user usersDomain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.conn.Exec(
		ctx,
		`INSERT INTO users (uid, name, age, email, password, role, balance) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		user.UID, user.Name, user.Age, user.Email,
		user.Password, user.Role, user.Balance,
	)
	if err != nil {
		// TODO: уточнение ошибки ( пользователь уже существует )
		return err
	}

	return nil
}

func (s *Storage) GetUserByEmail(email string) (usersDomain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := s.conn.QueryRow(ctx, "SELECT * FROM users WHERE email=$1", email)

	var user usersDomain.User
	if err := row.Scan(&user.UID, &user.Name, &user.Age, &user.Email, &user.Password, &user.Role, &user.Balance); err != nil {
		// TODO: уточнение ошибки ( нет такого пользователя )
		return usersDomain.User{}, err
	}

	return usersDomain.User{}, nil
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

package db

import (
	"context"
	"time"

	carsDomain "github.com/Dorrrke/rent-group1602/internal/domain/cars"
)

func (s *Storage) GetRentByRID(rid string) (carsDomain.Rent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := s.conn.QueryRow(ctx, "SELECT * FROM rents WHERE rid=$1", rid)

	var rent carsDomain.Rent

	if err := row.Scan(
		&rent.RID, &rent.CID, &rent.UID,
		&rent.Hours, &rent.Ended,
	); err != nil {
		return carsDomain.Rent{}, err
	}

	return rent, nil
}

func (s *Storage) StartRent(rent carsDomain.Rent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.conn.Exec(ctx,
		`INSERT INTO rents (car_id, user_id, hours, ended) 
		VALUES ($1, $2, $3, $4)`,
		rent.CID, rent.UID,
		rent.Hours, rent.Ended,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) EndRent(rid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.conn.Exec(ctx, "UPDATE rents SET ended=true WHERE rid=$1", rid)

	return err
}

func (s *Storage) GetRentHistoryByID(uid string) ([]carsDomain.Rent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.conn.Query(
		ctx,
		"SELECT * FROM rents WHERE user_id=$1",
		uid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rentHistory []carsDomain.Rent
	for rows.Next() {
		var rent carsDomain.Rent
		if err := rows.Scan(
			&rent.RID, &rent.CID, &rent.UID,
			&rent.Hours, &rent.Ended,
		); err != nil {
			return nil, err
		}

		rentHistory = append(rentHistory, rent)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rentHistory, nil
}

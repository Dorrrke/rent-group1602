package db

import (
	"context"
	"time"

	carsDomain "github.com/Dorrrke/rent-group1602/internal/domain/cars"
)

func (s *Storage) AddCar(car carsDomain.Car) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.conn.Exec(
		ctx,
		`INSERT INTO cars (cid, brand, model, color, year, number, price, available) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		car.CID, car.Brand, car.Model, car.Color,
		car.Year, car.Number, car.Price, car.Available,
	)
	if err != nil {
		// TODO: уточнение ошибки ( автомобиль уже существует )
		return err
	}

	return nil
}

func (s *Storage) GetAllCars() ([]carsDomain.Car, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.conn.Query(ctx, "SELECT * FROM cars WHERE available=true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []carsDomain.Car
	for rows.Next() {
		var car carsDomain.Car

		if err := rows.Scan(
			&car.CID, &car.Brand, &car.Model,
			&car.Color, &car.Year, &car.Number,
			&car.Price, &car.Available,
		); err != nil {
			return nil, err
		}

		cars = append(cars, car)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}

func (s *Storage) GetCarByCID(cid string) (carsDomain.Car, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := s.conn.QueryRow(ctx, "SELECT * FROM cars WHERE cid=$1", cid)

	var car carsDomain.Car
	if err := row.Scan(
		&car.CID, &car.Brand, &car.Model,
		&car.Color, &car.Year, &car.Number,
		&car.Price, &car.Available,
	); err != nil {
		//TODO: уточнение ошибки ( автомобиль не найден )
		return carsDomain.Car{}, err
	}

	return car, nil
}

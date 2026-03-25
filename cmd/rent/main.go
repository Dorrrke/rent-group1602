package main

import (
	"github.com/Dorrrke/rent-group1602/internal/repository/db"
	"github.com/Dorrrke/rent-group1602/internal/server"
	"github.com/Dorrrke/rent-group1602/internal/service/cars"
	"github.com/Dorrrke/rent-group1602/internal/service/profile"
	"github.com/Dorrrke/rent-group1602/internal/service/users"
)

func main() {
	dbDSN := "postgres://user:password@localhost:5432/rents?sslmode=disable"
	repo, err := db.New(dbDSN)
	if err != nil {
		panic(err)
	}
	if err := db.RunMigrations(dbDSN); err != nil {
		panic(err)
	}

	usersService := users.New(repo)
	carService := cars.New(repo)
	profileService := profile.New(repo)

	srv := server.New(":8080", usersService, carService, profileService)

	if err := srv.Run(); err != nil {
		panic(err)
	}
}

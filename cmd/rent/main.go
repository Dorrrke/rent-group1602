package main

import (
	"fmt"

	"github.com/Dorrrke/rent-group1602/internal"
	"github.com/Dorrrke/rent-group1602/internal/repository/db"
	"github.com/Dorrrke/rent-group1602/internal/server"
	"github.com/Dorrrke/rent-group1602/internal/service/cars"
	"github.com/Dorrrke/rent-group1602/internal/service/profile"
	"github.com/Dorrrke/rent-group1602/internal/service/users"
)

func main() {
	cfg := internal.ReadConfig()
	repo, err := db.New(cfg.DBDSN)
	if err != nil {
		panic(err)
	}
	if err := db.RunMigrations(cfg.DBDSN); err != nil {
		panic(err)
	}

	usersService := users.New(repo)
	carService := cars.New(repo)
	profileService := profile.New(repo)

	srv := server.New(
		fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		usersService,
		carService,
		profileService,
	)

	if err := srv.Run(); err != nil {
		panic(err)
	}
}

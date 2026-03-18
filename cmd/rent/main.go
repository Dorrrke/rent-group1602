package main

import (
	"github.com/Dorrrke/rent-group1602/internal/repository/memstorage"
	"github.com/Dorrrke/rent-group1602/internal/server"
	"github.com/Dorrrke/rent-group1602/internal/service/cars"
	"github.com/Dorrrke/rent-group1602/internal/service/users"
)

func main() {
	repo := memstorage.New()
	usersService := users.New(repo)
	carService := cars.New(repo)

	srv := server.New(":8080", usersService, carService)

	if err := srv.Run(); err != nil {
		panic(err)
	}
}

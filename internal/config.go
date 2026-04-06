package internal

import (
	"cmp"
	"os"
)

type Config struct {
	Host  string
	Port  string
	DBDSN string
}

// TODO: сделать валидацию кофига + конфигурация флагами

func ReadConfig() *Config {
	host := cmp.Or(os.Getenv("RENTAL_SERVICE_HOST"), "0.0.0.0")
	port := cmp.Or(os.Getenv("RENTAL_SERVICE_PORT"), "8080")
	dbdsn := cmp.Or(os.Getenv("RENTAL_SERVICE_DB_DSN"), "postgres://user:password@localhost:5432/rents?sslmode=disable")

	return &Config{
		Host:  host,
		Port:  port,
		DBDSN: dbdsn,
	}
}

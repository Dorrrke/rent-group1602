package internal

import (
	"cmp"
	"flag"
	"os"

	"github.com/rs/zerolog"
)

type Config struct {
	Host  string
	Port  string
	DBDSN string
	Debug bool
}

// TODO: сделать валидацию кофига + конфигурация флагами

func ReadConfig() *Config {
	var cfg Config
	flag.StringVar(&cfg.Host, "host", "0.0.0.0", "флаг для указания хоста для запуска сервиса")
	flag.StringVar(&cfg.Port, "port", "8080", "флаг для указания порта для запуска сервиса")
	flag.StringVar(&cfg.DBDSN, "dbdsn", "postgres://user:password@localhost:5432/rents?sslmode=disable", "флаг для указания строки подключения к БД")
	flag.BoolVar(&cfg.Debug, "debug", false, "флаг для запуска сервиса в режиме debug")
	flag.Parse()

	cfg.Host = cmp.Or(os.Getenv("RENTAL_SERVICE_HOST"), cfg.Host)
	cfg.Port = cmp.Or(os.Getenv("RENTAL_SERVICE_PORT"), cfg.Port)
	cfg.DBDSN = cmp.Or(os.Getenv("RENTAL_SERVICE_DB_DSN"), cfg.DBDSN)

	return &cfg
}

func (cfb *Config) ConfigureLogger() {
	if cfb.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		return
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

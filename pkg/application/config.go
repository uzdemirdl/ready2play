package application

import "os"

type Config struct {
	Port string
	DNS  string
}

func ReadConfigFromENV() (Config, error) {
	port := os.Getenv("APP_PORT")
	if len(port) == 0 {
		port = ":8080"
	}

	pgDNS := os.Getenv("DB_DSN")
	if len(pgDNS) == 0 {
		pgDNS = "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"
	}

	return Config{
		Port: port,
		DNS:  pgDNS,
	}, nil
}

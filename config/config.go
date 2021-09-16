package config

import (
	"errors"
	"os"
	"time"
)

type Config struct {
	ApiPort string
	DiscountClientHost string
	DiscountClientPort string
	BlackFridayDate time.Time
	DbFile string
}

func RetrieveConfig() (Config, error) {
	apiPort := os.Getenv("API_PORT")
	dbFile := os.Getenv("DB_FILE")
	blackFridayDate := os.Getenv("BLACK_FRIDAY_DATE")
	discountHost := os.Getenv("CLIENT_HOST")
	discountPort := os.Getenv("CLIENT_PORT")

	if len(apiPort) == 0 ||
		len(dbFile) == 0 ||
		len(blackFridayDate) == 0 ||
		len(discountHost) == 0 ||
		len(discountPort) == 0 {
		return Config{}, errors.New("Env error")
	}

	date, _ := time.Parse("2006-01-02", blackFridayDate)

	return Config{
		ApiPort: apiPort,
		DiscountClientHost: discountHost,
		DiscountClientPort: discountPort,
		BlackFridayDate: date,
		DbFile: dbFile,
	}, nil
}
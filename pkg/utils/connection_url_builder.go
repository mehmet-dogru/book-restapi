package utils

import (
	"fmt"
	"os"
)

func ConnectionUrlBuilder(n string) (string, error) {
	var url string

	switch n {
	case "postgres":
		//URL for PostgreSQL connection
		url = fmt.Sprintf(
			"host=%s post=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_SSL_MODE"),
		)
	case "mysql":
		//URL for MySQL connection
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_POST"),
			os.Getenv("DB_NAME"),
		)
	case "redis":
		//URL for Redis connection
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		)
	case "fiber":
		//URL for Fiber connection
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	return url, nil
}

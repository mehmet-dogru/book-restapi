package database

import (
	"book-restapi/pkg/utils"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"os"
	"strconv"
	"time"
)

func PostgresSQLConnection() (*sqlx.DB, error) {
	//Define database connection settings
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	//Build PostgreSQL connection URL
	postgresConnURL, err := utils.ConnectionUrlBuilder("postgres")
	if err != nil {
		return nil, err
	}

	//Define database connection for PostgreSQL
	db, err := sqlx.Connect("pgx", postgresConnURL)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, '%w'", err)
	}

	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	//Try to ping database
	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, not sent ping to database, '%w'", err)
	}

	return db, nil
}

package database

import (
	"book-restapi/app/queries"
	"github.com/jmoiron/sqlx"
	"os"
)

type Queries struct {
	*queries.UserQueries
	*queries.BookQueries
}

func NewDbConn(bookQueries *queries.BookQueries) *Queries {
	return &Queries{nil, bookQueries}
}

func OpenDBConnection() (*Queries, error) {
	//Define database connection variables
	var (
		db  *sqlx.DB
		err error
	)

	dbType := os.Getenv("DB_TYPE")

	//Define a new Database connection with right DB type
	switch dbType {
	case "pgx":
		db, err = PostgresSQLConnection()
	}

	if err != nil {
		return nil, err
	}

	return &Queries{
		//Set queries form models
		UserQueries: &queries.UserQueries{DB: db},
		BookQueries: &queries.BookQueries{DB: db},
	}, nil
}

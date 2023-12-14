package queries

import (
	"book-restapi/app/models"
	"github.com/jmoiron/sqlx"
)

type BookQueries struct {
	*sqlx.DB
}

// GetBooks method for getting all books
func (q *BookQueries) GetBooks() ([]models.Book, error) {
	//Define book variable
	var books []models.Book

	//Define query string
	query := `SELECT * FROM books`

	//Send query to database
	err := q.Select(&books, query)
	if err != nil {
		return books, err
	}

	//Return query result
	return books, nil
}

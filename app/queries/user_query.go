package queries

import (
	"book-restapi/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// UserQueries struct for queries form User Model
type UserQueries struct {
	*sqlx.DB
}

func (q *UserQueries) GetUserByID(id uuid.UUID) (models.User, error) {
	var user models.User

	query := `SELECT * FROM users WHERE id = $1`

	err := q.Get(&user, query, id)
	if err != nil {
		return user, err
	}

	return user, nil
}

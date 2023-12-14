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

func (q *UserQueries) CreateUser(u *models.User) error {
	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Send query to database.
	_, err := q.Exec(
		query,
		u.ID, u.CreatedAt, u.UpdatedAt, u.Email, u.PasswordHash, u.UserStatus, u.UserRole,
	)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

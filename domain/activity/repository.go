package activity

import (
	"database/sql"

	"github.com/jadahbakar/skyshi-todolist/util/logger"
)

type repo struct {
	db *sql.DB
}

type Repository interface {
	Create(*PostReq) (*Activity, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repo{db}
}

func (r *repo) Create(req *PostReq) (*Activity, error) {
	var t Activity
	query := `INSERT INTO activity(title, email, created_at, updated_at) VALUES ($1, $2, now(), now())
		RETURNING activity_id, title, email, created_at, updated_at`
	err := r.db.QueryRow(query, req.Title, req.Email).Scan(&t.Id, &t.Title, &t.Email, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		logger.Errorf("error: %v", err)
		return nil, err
	}
	return &t, nil
}

package activity

import (
	"fmt"

	"github.com/jadahbakar/skyshi-todolist/util/logger"
	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

type Repository interface {
	Create(*PostReq) (int64, error)
	Update(int64, string) (int64, error)
	Delete(int64) (int64, error)
	GetById(int64) (*Activity, error)
	GetAll() ([]Activity, error)
}

func NewRepository(db *sqlx.DB) Repository {
	return &repo{db}
}

func (r *repo) Create(req *PostReq) (int64, error) {
	query := fmt.Sprintf("INSERT INTO activities(title, email, created_at, updated_at) VALUES ('%s', '%s', now(), now())", req.Title, req.Email)
	res, err := r.db.Exec(query)
	if err != nil {
		logger.Errorf("error: %v", err)
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		logger.Errorf("last_insert_id: %v", err)
		return 0, err
	}

	return lastId, nil
}

func (r *repo) Update(id int64, title string) (int64, error) {
	query := fmt.Sprintf("UPDATE activities SET title = '%s', updated_at =  now() WHERE activity_id = %d", title, id)
	res, err := r.db.Exec(query)
	if err != nil {
		logger.Errorf("error: %v", err)
		return 0, err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		logger.Errorf("last_insert_id: %v", err)
		return 0, err
	}

	if affectedRows == 0 {
		logger.Errorf("no row affected: %v", err)
		return 0, err
	}

	return id, nil
}

func (r *repo) Delete(id int64) (int64, error) {
	query := fmt.Sprintf("DELETE FROM activities WHERE activity_id = %d", id)
	res, err := r.db.Exec(query)
	if err != nil {
		logger.Errorf("error: %v", err)
		return 0, err
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		logger.Errorf("last_insert_id: %v", err)
		return 0, err
	}

	if affectedRows == 0 {
		logger.Errorf("no row affected: %v", err)
		return 0, err
	}

	return id, nil
}

func (r *repo) GetById(id int64) (*Activity, error) {
	var t Activity
	query := "SELECT activity_id, title, email, created_at, updated_at FROM activities WHERE activity_id = ?"
	err := r.db.QueryRow(query, id).Scan(&t.Id, &t.Title, &t.Email, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		logger.Errorf("error: %v", err)
		return nil, err
	}
	return &t, nil
}

func (r *repo) GetAll() ([]Activity, error) {
	result := make([]Activity, 0)
	t := Activity{}
	query := "SELECT activity_id, title, email, created_at, updated_at FROM activities"
	rows, err := r.db.Query(query)
	if err != nil {
		logger.Errorf("Error Query: %v", err)
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(
			&t.Id,
			&t.Title,
			&t.Email,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	if rows.Err() != nil {
		logger.Errorf("Error Reading Rows: \n", err)
		return nil, rows.Err()
	}
	return result, nil
}
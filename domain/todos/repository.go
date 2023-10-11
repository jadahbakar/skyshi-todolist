package todos

import (
	"fmt"

	"github.com/jadahbakar/skyshi-todolist/util/logger"
	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

type Repository interface {
	Create(*PostReq) (int, error)
	Update(int, *PatchReq) (int, error)
	Delete(int) (int, error)
	GetById(int) (*Todo, error)
	GetAll(int) ([]Todo, error)
}

func NewRepository(db *sqlx.DB) Repository {
	return &repo{db}
}

func (r *repo) Create(req *PostReq) (int, error) {
	query := fmt.Sprintf("INSERT INTO todos(title, activity_group_id, is_active, priority, created_at, updated_at) VALUES ('%s', %d, %t, 'very-high',now(), now())", req.Title, req.ActivityId, req.IsActive)
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

	logger.Infof("LastInsertId: %d", int(lastId))

	return int(lastId), nil
}

func (r *repo) GetById(id int) (*Todo, error) {
	var t Todo
	query := "SELECT todo_id, title, activity_group_id, is_active, priority, created_at, updated_at FROM todos WHERE todo_id = ?"
	err := r.db.QueryRow(query, id).Scan(&t.Id, &t.Title, &t.ActivityId, &t.IsActive, &t.Priority, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		logger.Errorf("error: %v", err)
		return nil, err
	}
	return &t, nil
}

func (r *repo) Update(id int, req *PatchReq) (int, error) {
	query := fmt.Sprintf("UPDATE todos SET title = '%s', priority = '%s',  is_active = %t, updated_at = now() WHERE todo_id = %d", req.Title, req.Priority, req.IsActive, id)
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

func (r *repo) Delete(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM todos WHERE todo_id = %d", id)
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

func (r *repo) GetAll(id int) ([]Todo, error) {
	result := make([]Todo, 0)
	t := Todo{}
	query := fmt.Sprintf("SELECT todo_id, activity_group_id, title, is_active, priority, created_at, updated_at FROM todos WHERE activity_group_id = %d", id)
	rows, err := r.db.Query(query)
	if err != nil {
		logger.Errorf("Error Query: %v", err)
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(
			&t.Id,
			&t.ActivityId,
			&t.Title,
			&t.IsActive,
			&t.Priority,
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

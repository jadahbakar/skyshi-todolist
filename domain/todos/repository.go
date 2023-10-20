package todos

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jadahbakar/skyshi-todolist/util/errorlib"
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
	// query := fmt.Sprintf("INSERT INTO todos(title, activity_group_id, is_active, priority, created_at, updated_at) VALUES ('%s', %d, %t, 'very-high',now(), now())", req.Title, req.ActivityId, req.IsActive)
	query := fmt.Sprintf("INSERT INTO todos(title, activity_group_id, is_active, priority, created_at, updated_at) VALUES ('%s', %d, %t, 'very-high',now(), now())", req.Title, req.ActivityId, true)
	res, err := r.db.Exec(query)
	if err != nil {
		logger.Errorf("error: %v", err)
		return 0, errorlib.WrapErr(err, errorlib.ErrorCodeUnknown, "insert activities")
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		logger.Errorf("last_insert_id: %v", err)
		return 0, errorlib.WrapErr(err, errorlib.ErrorCodeNotFound, "id not found")
	}

	return int(lastId), nil
}

func (r *repo) GetById(id int) (*Todo, error) {
	var t Todo
	query := "SELECT id, title, activity_group_id, is_active, priority, created_at, updated_at FROM todos WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&t.Id, &t.Title, &t.ActivityId, &t.IsActive, &t.Priority, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		logger.Errorf("error: %v", err)
		return nil, errorlib.WrapErr(nil, errorlib.ErrorCodeNotFound, "Todo with ID %d Not Found", id)
	}
	return &t, nil
}

func (r *repo) Update(id int, req *PatchReq) (int, error) {
	query := fmt.Sprintf("UPDATE todos SET title = '%s', priority = '%s', is_active = %t, status = '%s', updated_at = now() WHERE id = %d", req.Title, "very-high", req.IsActive, req.Status, id)
	log.Info(query)
	res, err := r.db.Exec(query)
	if err != nil {
		logger.Errorf("error: %v", err)
		return 0, errorlib.WrapErr(err, errorlib.ErrorCodeInternal, "update activities")
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		logger.Errorf("last_insert_id: %v", err)
		return 0, errorlib.WrapErr(err, errorlib.ErrorCodeInternal, "update activities")
	}

	if affectedRows == 0 {
		logger.Errorf("no row affected: %v", err)
		return 0, errorlib.WrapErr(err, errorlib.ErrorCodeInternal, "affected rows")
	}

	return id, nil
}

func (r *repo) Delete(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM todos WHERE id = %d", id)
	res, err := r.db.Exec(query)
	if err != nil {
		logger.Errorf("error: %v", err)
		return 0, errorlib.WrapErr(err, errorlib.ErrorCodeInternal, "delete activities")
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		logger.Errorf("last_insert_id: %v", err)
		return 0, errorlib.WrapErr(err, errorlib.ErrorCodeInternal, "affected rows")
	}

	if affectedRows == 0 {
		logger.Errorf("no row affected: %v", err)
		return 0, errorlib.WrapErr(err, errorlib.ErrorCodeInternal, "affected rows")
	}

	return id, nil
}

func (r *repo) GetAll(id int) ([]Todo, error) {
	result := make([]Todo, 0)
	t := Todo{}
	query := fmt.Sprintf("SELECT id, activity_group_id, title, is_active, priority, created_at, updated_at FROM todos WHERE activity_group_id = %d", id)
	rows, err := r.db.Query(query)
	if err != nil {
		logger.Errorf("Error Query: %v", err)
		return nil, errorlib.WrapErr(err, errorlib.ErrorCodeInvalidArgument, "error query")
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
			return nil, errorlib.WrapErr(err, errorlib.ErrorCodeInvalidArgument, "error scan")
		}
		result = append(result, t)
	}
	if rows.Err() != nil {
		logger.Errorf("Error Reading Rows: \n", err)
		return nil, errorlib.WrapErr(rows.Err(), errorlib.ErrorCodeInternal, "error reading rows")
	}
	return result, nil
}

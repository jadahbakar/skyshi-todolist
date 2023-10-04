package mysql

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jadahbakar/skyshi-todolist/util/config"
)

type mysqlRepo struct {
	db *sqlx.DB
}

func Connected(config *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", config.Db.Url)
	if err != nil {
		return nil, err
	}
	return db, nil
}

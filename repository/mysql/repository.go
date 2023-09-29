package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jadahbakar/skyshi-todolist/util/config"
)

type mysqlRepo struct {
	db *sql.DB
}

func NewMysqlClient(config *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Db.User,
		config.Db.Password,
		config.Db.Host,
		config.Db.Port,
		config.Db.Name,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

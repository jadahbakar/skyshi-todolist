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
	// fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	// 	config.Db.User,
	// 	config.Db.Password,
	// 	config.Db.Host,
	// 	config.Db.Port,
	// 	config.Db.Name,
	// )
	// log.Info(dsn)
	// db, err := sqlx.Connect("mysql", config.Db.Url)
	dsn := "todo:secret@tcp(mysql:3306)/todolist"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

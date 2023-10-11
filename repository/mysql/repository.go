package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jadahbakar/skyshi-todolist/util/config"
	"github.com/jadahbakar/skyshi-todolist/util/logger"
)

type mysqlRepo struct {
	db *sqlx.DB
}

func Connected(config *config.Config) (*sqlx.DB, error) {
	mysqlHost := config.Db.Host
	mysqlPort := config.Db.Port
	mysqlUser := config.Db.User
	mysqlPass := config.Db.Password
	mysqlDb := config.Db.Name

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDb)
	logger.Infof("dataSourceName: %v", dataSourceName)
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

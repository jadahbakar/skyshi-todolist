package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jadahbakar/skyshi-todolist/domain"
	"github.com/jadahbakar/skyshi-todolist/repository/mysql"
	"github.com/jadahbakar/skyshi-todolist/util/config"
	"github.com/jadahbakar/skyshi-todolist/util/engine"
	"github.com/jadahbakar/skyshi-todolist/util/logger"
)

func main() {
	log.Printf("Define Config......")
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error loading config:%v\n", err)
	}

	logger.NewAppLogger(config)

	log.Printf("Define Engine......")
	server, err := engine.NewFiber(config)
	if err != nil {
		log.Fatalf("error starting server:%v\n", err)
	}

	log.Printf("Connecting DB......")
	db, err := mysql.Connected(config)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("Connected to DB....")

	log.Printf("Prepare Migration..")

	mysqlHost := config.Db.Host
	mysqlPort := config.Db.Port
	mysqlUser := config.Db.User
	mysqlPass := config.Db.Password
	mysqlDb := config.Db.Name

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDb)

	driver, err := migrate.New("file://"+config.Db.MigrationFolder, "mysql://"+dsn)
	if err != nil {
		log.Fatalf("Error creating migration driver: %s", err)
	}
	if err := driver.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error applying migrations: %s", err)
	}
	log.Println("Migrations applied successfully!")

	domain.MonolithIOC(server, db)

	engine.StartFiberWithGracefulShutdown(server, config.App.Port)
}

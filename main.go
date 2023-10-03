package main

import (
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
	// viper.SetConfigFile("migrate.yml")
	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatalf("Error reading config file: %s", err)
	// }
	// driver, err := migrate.New(
	// 	"file://"+viper.GetString("database.migrations_folder"),
	// 	viper.GetString("database.connection_string"),
	// )
	// if err != nil {
	// 	log.Fatalf("Error creating migration driver: %s", err)
	// }

	// if err := driver.Up(); err != nil && err != migrate.ErrNoChange {
	// 	log.Fatalf("Error applying migrations: %s", err)
	// }
	driver, err := migrate.New("file://"+config.Db.MigrationFolder, "mysql://"+config.Db.Url)
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

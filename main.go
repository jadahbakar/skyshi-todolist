package main

import (
	"log"

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

	domain.MonolithIOC(server, db)

	engine.StartFiberWithGracefulShutdown(server, config.App.Port)
}

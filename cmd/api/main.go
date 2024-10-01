package main

import (
	"log"
	"os"

	"github.com/Kenini1805/go-rest-api/config"
	"github.com/Kenini1805/go-rest-api/internal/server"
	"github.com/Kenini1805/go-rest-api/pkg/db/postgres"
	"github.com/Kenini1805/go-rest-api/pkg/logger"
	"github.com/Kenini1805/go-rest-api/pkg/utils"
)

func main() {
	log.Println("Starting api server ...")
	configPath := utils.GetConfigPath(os.Getenv("config"))
	cfgFile, err := config.LoadConfig(configPath)

	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)

	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	appLogger := logger.NewAPILogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v",
		cfg.Server.AppVersion,
		cfg.Logger.Level,
		cfg.Server.Mode,
		cfg.Server.SSL,
	)

	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
	} else {
		appLogger.Infof("Postgres connected, Status: %#v", psqlDB.Stats())
	}
	defer psqlDB.Close()

	s := server.NewServer(cfg, psqlDB, appLogger)
	if err = s.Run(); err != nil {
		log.Println(err)
	}
}

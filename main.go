package main

import (
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/cli"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/config"
	"github.com/Joroboro253/ReviewApiDistributedLab/internal/service/repository"
	"log"
)

func main() {
	cfg, err := config.ReadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	db, err := repository.NewPostgresDB(cfg.DB.URL)
	if err != nil {
		log.Fatalf("Failed to initialize db: %v", err)
	}
	defer db.Close()

	// Initialization and starting application
	app := cli.NewApp(db)
	err = app.Start(cfg.Server.Port)
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

}

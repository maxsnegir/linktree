package main

import (
	"linktree/cmd/config"
	"linktree/cmd/logging"
	"linktree/intrernal/server"
	"log"
)

func main() {
	cfg, err := config.NewConfig("cfg.yaml")
	if err != nil {
		log.Fatal(err)
	}
	logger, err := logging.NewLogger(cfg.Logging.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	srv := server.NewServer(cfg, logger)
	logger.Fatal(srv.Start())
}

package main

import (
	"github.com/Masih-Ghasri/GolangBackend/api"
	"github.com/Masih-Ghasri/GolangBackend/config"
	"github.com/Masih-Ghasri/GolangBackend/data/cache"
	"github.com/Masih-Ghasri/GolangBackend/data/db"
	"github.com/Masih-Ghasri/GolangBackend/data/db/migrations"
	"github.com/Masih-Ghasri/GolangBackend/pkg/logging"
)

func main() {
	cfg := config.Getconfig()
	logger := logging.NewLogger(cfg)

	//Redis
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	//DataBase
	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}

	migrations.Up1()
	//BackEnd
	api.Initserver(cfg)
}

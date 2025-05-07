package db

import (
	"fmt"
	"github.com/Masih-Ghasri/GolangBackend/config"
	"github.com/Masih-Ghasri/GolangBackend/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var logger = logging.NewLogger(config.Getconfig())
var dbClient *gorm.DB

func InitDb(cfg *config.Config) error {
	var err error
	cnn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s ",
		cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Dbname, cfg.Postgres.SSLMode)

	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := dbClient.DB()
	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)

	sqlDB.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)
	logger.Info(logging.Postgres, logging.Startup, "Connected to postgres", nil)
	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	if dbClient == nil {
		logger.Warn(logging.Postgres, logging.Rollback, "Database client is nil, nothing to close", nil)
		return
	}

	con, err := dbClient.DB()
	if err != nil {
		logger.Error(logging.Postgres, logging.Rollback, "Failed to get sql.DB: "+err.Error(), nil)
		return
	}

	if err := con.Close(); err != nil {
		logger.Error(logging.Postgres, logging.Rollback, "Failed to close database: "+err.Error(), nil)
	} else {
		logger.Info(logging.Postgres, logging.Rollback, "Database connection closed", nil)
	}
}

package migrations

import (
	"github.com/Masih-Ghasri/GolangBackend/config"
	"github.com/Masih-Ghasri/GolangBackend/data/db"
	"github.com/Masih-Ghasri/GolangBackend/data/models"
	"github.com/Masih-Ghasri/GolangBackend/pkg/logging"
)

var logger = logging.NewLogger(config.Getconfig())

func Up1() {
	database := db.GetDb()
	if database == nil {
		logger.Fatal(logging.Postgres, logging.Migration, "Database client is nil", nil)
	}
	//"Masih-Ghasri change me" add some model
	err := database.AutoMigrate(&models.User{})
	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, "Failed to create User table", map[logging.ExtraKey]interface{}{
			"error": err.Error(),
		})
	}

	logger.Info(logging.Postgres, logging.Migration, "Database migration complete", nil)
}

func Down1() {}

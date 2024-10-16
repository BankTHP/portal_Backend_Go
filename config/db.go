package config

import (
	"fmt"
	"log"
	"pccth/portal-blog/internal/entity"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


func InitConfig() {
	viper.SetConfigName("config.dev")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

}

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.dbname"))

	var logMode logger.LogLevel
	if viper.GetBool("app.debug") {
		logMode = logger.Info
	} else {
		logMode = logger.Silent
	}

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	}
	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := autoMigrateEntities(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}

func autoMigrateEntities(db *gorm.DB) error {
	models := []interface{}{
		&entity.Post{},&entity.Comment{},&entity.News{},&entity.Notification{},&entity.Release{},
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return err
		}
	}
	return nil
}

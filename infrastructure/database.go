package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/takumi616/english-vocabularies-database/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(ctx context.Context, cfg *Config) (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.PassWord, cfg.DbName, cfg.Sslmode)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to open postgresql: %v", err)
		return nil, err
	}

	db.AutoMigrate(&entities.Vocabulary{})
	return db, nil
}

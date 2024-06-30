package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/takumi616/english-vocabularies-database/entities"
	"github.com/takumi616/english-vocabularies-database/infrastructure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Db *gorm.DB
}

func Init(ctx context.Context, cfg *infrastructure.Config) (*Postgres, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.PassWord, cfg.DbName, cfg.Sslmode)
	log.Printf("db info: %s", dataSourceName)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to open postgresql: %v", err)
		return nil, err
	}

	db.AutoMigrate(&entities.Vocabulary{})

	postgres := &Postgres{Db: db}
	return postgres, nil
}

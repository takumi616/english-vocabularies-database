package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/takumi616/english-vocabularies-database/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	PgConfig *PgConfig
}

func (p *Postgres) InitDatabase(ctx context.Context) (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", p.PgConfig.Host, p.PgConfig.Port, p.PgConfig.User, p.PgConfig.PassWord, p.PgConfig.DbName, p.PgConfig.Sslmode)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to open postgresql: %v", err)
		return nil, err
	}

	db.AutoMigrate(&entities.Vocabulary{})
	return db, nil
}

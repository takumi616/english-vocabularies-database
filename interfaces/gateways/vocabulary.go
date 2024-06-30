package gateways

import (
	"context"

	"github.com/takumi616/english-vocabularies-database/entities"
	"gorm.io/gorm"
)

type VocabularyGateway struct {
	Db *gorm.DB
}

func NewVocabularyGateway(db *gorm.DB) *VocabularyGateway {
	return &VocabularyGateway{
		Db: db,
	}
}

func (vg *VocabularyGateway) AddNewVocabulary(ctx context.Context, vocabulary *entities.Vocabulary) (int64, error) {
	//DB proccess

	return 1, nil
}

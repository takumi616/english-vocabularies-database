package gateways

import (
	"context"

	"github.com/takumi616/english-vocabularies-database/entities"
)

type VocabularyGateway struct {
	Postgres Postgres
}

func (vg *VocabularyGateway) AddNewVocabulary(ctx context.Context, vocabulary *entities.Vocabulary) (int64, error) {
	db, _ := vg.Postgres.InitDatabase(ctx)

	tx := db.Create(vocabulary)
	tx.Commit()
	return 1, nil
}

package repositories

import (
	"context"

	"github.com/takumi616/english-vocabularies-database/entities"
)

type VocabularyRepository interface {
	AddNewVocabulary(ctx context.Context, vocabulary *entities.Vocabulary) (int64, error)
}

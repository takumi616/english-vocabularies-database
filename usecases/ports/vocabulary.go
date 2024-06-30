package ports

import (
	"context"

	"github.com/takumi616/english-vocabularies-database/entities"
)

type VocabularyInputPort interface {
	AddNewVocabulary(ctx context.Context, vocabulary *entities.Vocabulary) error
}

type VocabularyOutputPort interface {
	WriteResponse(vocabularyID int64) error
	WriteError(errMsg error) error
}

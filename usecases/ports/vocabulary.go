package ports

import (
	"context"
	"net/http"

	"github.com/takumi616/english-vocabularies-database/entities"
)

type VocabularyInputPort interface {
	AddNewVocabulary(ctx context.Context, w http.ResponseWriter, vocabulary *entities.Vocabulary) error
}

type VocabularyOutputPort interface {
	WriteResponse(ctx context.Context, w http.ResponseWriter, vocabularyID int64) error
	WriteError(ctx context.Context, w http.ResponseWriter, errMsg error) error
}

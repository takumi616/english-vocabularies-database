package repositories

import "github.com/takumi616/english-vocabularies-database/entities"

type VocabularyRepository interface {
	AddNewVocabulary(vocabulary *entities.Vocabulary) (int64, error)
}

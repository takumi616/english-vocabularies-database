package persistences

import "github.com/takumi616/english-vocabularies-database/entities"

type VocabularyPersistence interface {
	Add(vocabulary *entities.Vocabulary) (int64, error)
}

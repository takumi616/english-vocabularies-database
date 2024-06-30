package gateways

import (
	"github.com/takumi616/english-vocabularies-database/entities"
	"github.com/takumi616/english-vocabularies-database/interfaces/persistences"
)

type VocabularyGateway struct {
	//interface to infrastructure DB
	Persistence persistences.VocabularyPersistence
}

func (vg *VocabularyGateway) AddNewVocabulary(vocabulary *entities.Vocabulary) (int64, error) {
	return vg.Persistence.Add(vocabulary)
}

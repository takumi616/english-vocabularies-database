package interactors

import (
	"github.com/takumi616/english-vocabularies-database/entities"
	"github.com/takumi616/english-vocabularies-database/usecases/repositories"
)

type VocabularyInteractor struct {
	//repository interface
	Repo repositories.VocabularyRepository
	//output ports interface
}

func (vi *VocabularyInteractor) AddNewVocabulary(title string, example string) (int64, error) {
	//Call repository interfaces
	vocabulary := entities.NewVocabulary(title, example)
	return vi.Repo.AddNewVocabulary(vocabulary)
}

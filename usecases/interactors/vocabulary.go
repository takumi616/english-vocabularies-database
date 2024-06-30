package interactors

import (
	"context"
	"net/http"

	"github.com/takumi616/english-vocabularies-database/entities"
	"github.com/takumi616/english-vocabularies-database/interfaces/gateways"
	"github.com/takumi616/english-vocabularies-database/interfaces/presenters"
	"github.com/takumi616/english-vocabularies-database/usecases/ports"
	"github.com/takumi616/english-vocabularies-database/usecases/repositories"
)

type VocabularyInteractor struct {
	//repository interface
	Repo repositories.VocabularyRepository

	//output ports interface
	OutputPort ports.VocabularyOutputPort
}

func NewVocabularyInteractor(gateway *gateways.VocabularyGateway, presenter *presenters.VocabularyPresenter) *VocabularyInteractor {
	return &VocabularyInteractor{
		Repo:       gateway,
		OutputPort: presenter,
	}
}

func (vi *VocabularyInteractor) AddNewVocabulary(ctx context.Context, w http.ResponseWriter, vocabulary *entities.Vocabulary) error {
	addedID, _ := vi.Repo.AddNewVocabulary(ctx, vocabulary)
	//Call outputPort
	err := vi.OutputPort.WriteResponse(ctx, w, addedID)
	return err
}

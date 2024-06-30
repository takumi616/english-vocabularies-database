package presenters

import (
	"context"
	"encoding/json"
	"net/http"
)

type VocabularyPresenter struct {
}

func NewVocabularyPresenter() *VocabularyPresenter {
	return &VocabularyPresenter{}
}

func (vp *VocabularyPresenter) WriteResponse(ctx context.Context, w http.ResponseWriter, vocabularyID int64) error {
	rsp := struct {
		VocabularyID int64 `json:"vocabulary_id"`
	}{VocabularyID: vocabularyID}
	return json.NewEncoder(w).Encode(rsp)
}

func (vp *VocabularyPresenter) WriteError(ctx context.Context, w http.ResponseWriter, errMsg error) error {
	return nil
}

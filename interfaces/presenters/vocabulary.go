package presenters

type VocabularyPresenter struct {
}

func NewVocabularyPresenter() *VocabularyPresenter {
	return &VocabularyPresenter{}
}

type VocabularyOutputPort interface {
	WriteResponse(vocabularyID int64) error
	WriteError(errMsg error) error
}

func (vp *VocabularyPresenter) WriteResponse(vocabularyID int64) error {
	return nil
}

func (vp *VocabularyPresenter) WriteError(errMsg error) error {
	return nil
}

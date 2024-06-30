package inputPorts

type VocabularyInputPort interface {
	AddNewVocabulary(title string, example string) (int64, error)
}

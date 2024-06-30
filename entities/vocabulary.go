package entities

type Vocabulary struct {
	Title   string `json:"title"`
	Example string `json:"example"`
}

func NewVocabulary(title string, example string) *Vocabulary {
	return &Vocabulary{
		Title:   title,
		Example: example,
	}
}

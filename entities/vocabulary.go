package entities

import "gorm.io/gorm"

type Vocabulary struct {
	gorm.Model
	Title   string `json:"title"`
	Example string `json:"example"`
}

func NewVocabulary(title string, example string) *Vocabulary {
	return &Vocabulary{
		Title:   title,
		Example: example,
	}
}

package postgres

import (
	"github.com/takumi616/english-vocabularies-database/entities"
)

type Postgres struct {
	//dummy
	DbHandle string
}

func (p *Postgres) Add(vocabulary *entities.Vocabulary) (int64, error) {
	//DB proccess

	return 1, nil
}

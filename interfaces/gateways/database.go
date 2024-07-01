package gateways

import (
	"context"

	"gorm.io/gorm"
)

type Postgres interface {
	InitDatabase(ctx context.Context) (*gorm.DB, error)
}

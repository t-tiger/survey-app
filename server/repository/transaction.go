package repository

import (
	"context"

	"gorm.io/gorm"
)

type Transactional interface {
	WithTransaction(ctx context.Context, repos []Transactional, exec func(repos []Transactional) error) error
	StartTransaction(db *gorm.DB) Transactional
}

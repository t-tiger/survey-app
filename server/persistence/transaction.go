package persistence

import (
	"context"

	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/repository"
	"gorm.io/gorm"
)

type transaction struct {
	db *gorm.DB
}

func (t *transaction) WithTransaction(
	ctx context.Context, repos []repository.Transactional, exec func(repos []repository.Transactional) error,
) error {
	db := t.db.WithContext(ctx).Begin()

	newRepos := make([]repository.Transactional, len(repos))
	for i, repo := range repos {
		newRepos[i] = repo.StartTransaction(db)
	}
	err := exec(newRepos)
	if err != nil {
		if rollbackErr := db.Rollback().Error; rollbackErr != nil {
			return cerrors.Errorf(cerrors.DatabaseErr, rollbackErr.Error())
		}
		return cerrors.Errorf(cerrors.DatabaseErr, err.Error())
	}
	if err = db.Commit().Error; err != nil {
		return cerrors.Errorf(cerrors.DatabaseErr, err.Error())
	}
	return nil
}

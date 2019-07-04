package usecase

import (
	"context"

	"bitbucket.org/kamiazya/tcho2/contexts/tag-manager/domain/model"
	"bitbucket.org/kamiazya/tcho2/contexts/tag-manager/domain/model/tag"
	"bitbucket.org/kamiazya/tcho2/contexts/tag-manager/domain/repository/tagsrepo"
)

func NewStoreTagUsecase(repo tagsrepo.Repository) StoreTagUsecase {
	return &storeTagUsecase{
		repo: repo,
	}
}

type StoreTagUsecase interface {
	StoreTag(ctx context.Context, t *tag.Tag) (ID model.ID, err error)
}

type storeTagUsecase struct {
	repo tagsrepo.Repository
}

func (uc *storeTagUsecase) StoreTag(ctx context.Context, t *tag.Tag) (ID model.ID, err error) {
	ID, err = uc.repo.Store(ctx, t)
	return
}

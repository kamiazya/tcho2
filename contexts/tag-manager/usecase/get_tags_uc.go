package usecase

import (
	"context"

	"bitbucket.org/kamiazya/tcho2/contexts/tag-manager/domain/model/tag"
	"bitbucket.org/kamiazya/tcho2/contexts/tag-manager/domain/repository/tagsrepo"
)

func NewGetTagsUsecase(repo tagsrepo.Repository) GetTagsUsecase {
	return &getTagsUsecase{
		repo: repo,
	}
}

type GetTagsUsecase interface {
	GetTags(ctx context.Context, opts ...tagsrepo.Option) (tags []tag.Tag, err error)
}

type getTagsUsecase struct {
	repo tagsrepo.Repository
}

func (uc *getTagsUsecase) GetTags(ctx context.Context, opts ...tagsrepo.Option) (tags []tag.Tag, err error) {
	tags, err = uc.repo.Find(ctx, opts...)
	return
}

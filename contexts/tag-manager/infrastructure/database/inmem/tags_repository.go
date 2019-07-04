package inmem

import (
	"context"

	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/model"
	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/model/tag"
	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/repository"
	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/repository/tagsrepo"
)

func NewTagsRepository(opts ...TagsRepositoryOption) (repo tagsrepo.Repository, err error) {
	config := DefaultConfig()
	if len(opts) != 0 {
		if err := config.apply(opts); err != nil {
			return nil, err
		}
	}
	return NewTagsRepositoryFromConfig(config)
}

func NewTagsRepositoryFromConfig(config *TagsRepositoryConfig) (repo tagsrepo.Repository, err error) {
	return config.build()
}

func (repo *tagsRepository) Store(ctx context.Context, t *tag.Tag) (model.ID, error) {
	// IDをすでに持っている
	if t.ID != nil && !t.ID.IsNew() {
		return nil, repository.ErrAleadyHasID
	}

	// 名前が空文字になっている
	if t.Name == "" {
		return nil, tagsrepo.ErrTagNameIsRequired
	}

	// ロック
	repo.Lock()
	defer repo.Unlock()

	return repo.store(ctx, t)
}

func (repo *tagsRepository) Update(ctx context.Context, ID model.ID, t *tag.Tag) error {

	if ID != nil && ID.IsNew() {
		return repository.ErrNotHaveID
	}

	// 名前が空文字になっている
	if t.Name == "" {
		return tagsrepo.ErrTagNameIsRequired
	}

	if !repo.exists(ID.String()) {
		return repository.ErrNotFound
	}

	repo.Lock()
	defer repo.Unlock()
	return repo.update(ctx, ID, t)
}

func (repo *tagsRepository) Find(ctx context.Context, opts ...tagsrepo.Option) (tags []tag.Tag, err error) {
	repo.Lock()
	defer repo.Unlock()
	return repo.find(ctx, opts)
}

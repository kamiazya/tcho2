package inmem

import (
	"context"

	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/model/tag"
	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/repository/tagsrepo"
)

func DefaultConfig() *TagsRepositoryConfig {
	return &TagsRepositoryConfig{
		InitialCap: TAGS_REPO_DEFAULT_CAP,
	}
}

type TagsRepositoryConfig struct {
	InitialCap int
	WithTags   []tag.Tag
}

func (config *TagsRepositoryConfig) apply(opts []TagsRepositoryOption) error {
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return err
		}
	}
	return nil
}

func (config *TagsRepositoryConfig) build() (tagsrepo.Repository, error) {
	repo := new(tagsRepository)
	if config.WithTags != nil && config.InitialCap < len(config.WithTags) {
		config.InitialCap = len(config.WithTags)
	}

	repo.tags = make([]tag.Tag, 0, config.InitialCap)
	repo.idmap = make(map[string]struct{}, config.InitialCap)

	if config.WithTags != nil {
		ctx := context.TODO()
		for _, t := range config.WithTags {
			repo.store(ctx, &t)
		}
	}
	return repo, nil
}

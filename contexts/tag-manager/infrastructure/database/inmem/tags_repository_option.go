package inmem

import "bitbucket.org/kamiazya/tcho2/contexts/tag-manager/domain/model/tag"

type TagsRepositoryOption func(*TagsRepositoryConfig) error

func TagsRepoInitialCap(cap int) TagsRepositoryOption {
	return func(config *TagsRepositoryConfig) error {
		if config.InitialCap != TAGS_REPO_DEFAULT_CAP {
			return ErrInvalidRepoOption
		}
		config.InitialCap = cap
		return nil
	}
}

func TagsRepoWithTags(tags []tag.Tag) TagsRepositoryOption {
	return func(config *TagsRepositoryConfig) error {
		if config.WithTags != nil {
			return ErrInvalidRepoOption
		}
		config.WithTags = tags
		return nil
	}
}

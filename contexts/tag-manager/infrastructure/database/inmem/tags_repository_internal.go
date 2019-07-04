package inmem

import (
	"context"
	"sync"

	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/model"
	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/model/tag"
	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/repository/tagsrepo"
	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/values/stringid"
	"github.com/rs/xid"
)

type tagsRepository struct {
	sync.Mutex
	tags  []tag.Tag
	idmap map[string]struct{}
}

func (repo *tagsRepository) newID() model.ID {
	var (
		newIDCreated bool
		idStr        string
	)
	for !newIDCreated {
		idStr = xid.New().String()
		if _, exists := repo.idmap[idStr]; !exists {
			repo.idmap[idStr] = struct{}{}
			newIDCreated = true
		}
	}
	return stringid.ID(idStr)
}

func (repo *tagsRepository) store(ctx context.Context, t *tag.Tag) (ID model.ID, err error) {
	ID = repo.newID()
	t.ID = ID
	repo.tags = append(repo.tags, *t)
	return
}

func (repo *tagsRepository) find(ctx context.Context, opts []tagsrepo.Option) (tags []tag.Tag, err error) {
	fs, err := newFilters(opts)
	if err != nil {
		return nil, err
	}

	tags = make([]tag.Tag, 0, len(repo.tags))
	for _, t := range repo.tags {
		if fs.match(t) {
			tags = append(tags, t)
		}
	}
	return
}

func (repo *tagsRepository) exists(idStr string) (exists bool) {
	_, exists = repo.idmap[idStr]
	return
}

func (repo *tagsRepository) update(ctx context.Context, ID model.ID, t *tag.Tag) error {
	t.ID = ID
	for i, ta := range repo.tags {
		if tag.IsSame(ta, *t) {
			repo.tags[i] = *t
		}
	}
	return nil
}

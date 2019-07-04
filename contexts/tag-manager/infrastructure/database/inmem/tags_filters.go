package inmem

import (
	"bitbucket.org/kamiazya/tcho2/contexts/tag-manager/domain/model/tag"
	"bitbucket.org/kamiazya/tcho2/contexts/tag-manager/domain/repository/tagsrepo"
)

func newFilters(opts tagsrepo.Options) (fs filters, err error) {
	stmp, err := opts.Stmp()
	if err != nil {
		return nil, err
	}

	fs = make([]filter, 0, 1)
	if stmp.ByID != nil && !stmp.ByID.IsNew() {
		fs = append(fs, byIDFilter(stmp.ByID))
	}
	return
}

type filters []filter

func (fs filters) match(t tag.Tag) bool {
	for _, f := range fs {
		if !f(t) {
			return false
		}
	}
	return true
}

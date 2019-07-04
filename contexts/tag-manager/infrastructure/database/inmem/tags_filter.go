package inmem

import (
	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/model"
	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/model/tag"
)

type filter func(t tag.Tag) bool

func byIDFilter(ID model.ID) filter {
	return func(t tag.Tag) bool {
		return t.ID.String() == ID.String()
	}
}

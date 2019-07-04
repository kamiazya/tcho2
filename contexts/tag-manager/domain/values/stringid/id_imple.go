package stringid

import "bitbucket.org/kamiazya/tcho2/contexts/tag-manager/domain/model"

type stringID string

var _ model.ID = (*stringID)(nil)

func (id *stringID) String() string {
	if id == nil {
		return ""
	}
	return string(*id)
}

func (id *stringID) IsNew() bool {
	if id == nil {
		return true
	}
	return id.String() == ""
}

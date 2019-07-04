package tagsrepo

import (
	"bitbucket.org/kamiazya/tcho2/contexts/tag-manager/domain/model"
	"bitbucket.org/kamiazya/tcho2/contexts/tag-manager/domain/repository"
)

// Option is a search condition of Tag for the repository.
//
// Option はリポジトリに対してのTagの検索条件。
type Option func(*Stmp) error

// ByID is an option when searching by ID for Tag repository.
//
// ByID はTagリポジトリに対してIDで検索を行うときのOption。
func ByID(ID model.ID) Option {
	return func(stmp *Stmp) error {
		if stmp.ByID != nil {
			return repository.ErrOptionAlreadySeted
		}

		// 新しいIDでは検索できない
		if ID.IsNew() {
			return repository.ErrInvalidSearchCondition
		}

		stmp.ByID = ID
		return nil
	}
}

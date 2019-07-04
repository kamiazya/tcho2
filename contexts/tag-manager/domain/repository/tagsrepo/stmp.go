package tagsrepo

import "github.com/kamiazya/tcho2/contexts/tag-manager/domain/model"

// Stmp is a search statement for the Tag repository.
//
// Stmp は、 Tag リポジトリに対しての検索ステートメント。
type Stmp struct {
	ByID model.ID
}

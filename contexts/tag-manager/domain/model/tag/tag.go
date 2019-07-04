package tag

import "github.com/kamiazya/tcho2/contexts/tag-manager/domain/model"

// Tag is a model representing information that can be given to multiple entities.
//
// Tagは、エンティティに対して複数付与できる情報を表すモデルである。
type Tag struct {
	model.ID
	Name string
}

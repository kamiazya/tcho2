package stringid

import (
	"bitbucket.org/kamiazya/tcho2/contexts/tag-manager/domain/model"
)

// NewID specifies a new ID.
//
// 新規IDを明示させる。
func NewID() model.ID {
	return (*stringID)(nil)
}

// ID returns the ID corresponding to the character string given as an argument.
//
// 引数で与えられた文字列に対応するIDを返す。
// 空文字列をIDとしようとしたとき、 新規のIDを作成する。
func ID(id string) model.ID {
	if id == "" {
		return NewID()
	}
	tID := stringID(id)
	return &tID
}

package tagsrepo

import (
	"testing"

	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/model"
	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/repository"
	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/values/stringid"
)

func TestByID(t *testing.T) {
	testcases := []struct {
		title     string
		stmp      *Stmp
		ID        model.ID
		exceptErr error
	}{
		{
			title:     "正常系",
			stmp:      new(Stmp),
			ID:        stringid.ID("1"),
			exceptErr: nil,
		},
		{
			title: "異常系 すでに同じ検索条件がセットされている",
			stmp: func() *Stmp {
				stmp := new(Stmp)
				opt := ByID(stringid.ID("2"))
				if err := opt(stmp); err != nil {
					// ここには入れないはずなのでFatal
					t.Fatal(err)
				}
				return stmp
			}(),
			ID:        stringid.ID("1"),
			exceptErr: repository.ErrOptionAlreadySeted,
		},
		{
			title:     "異常系 空文字列をセットする",
			stmp:      new(Stmp),
			ID:        stringid.ID(""),
			exceptErr: repository.ErrInvalidSearchCondition,
		},
		{
			title:     "異常系 nilをセットする",
			stmp:      new(Stmp),
			ID:        stringid.NewID(),
			exceptErr: repository.ErrInvalidSearchCondition,
		},
	}

	for _, tc := range testcases {
		// 検索 Option を作成
		opt := ByID(tc.ID)

		// stmpに条件を適応する
		if err := opt(tc.stmp); err != tc.exceptErr {
			t.Error(tc.title, err, tc.exceptErr)
		}
	}
}

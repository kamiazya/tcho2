package stringid

import (
	"testing"

	"bitbucket.org/kamiazya/tcho2/contexts/tag-manager/domain/model"
)

func TestID_String(t *testing.T) {

	testcases := []struct {
		ID     model.ID
		output string
	}{
		{
			ID:     ID("aaa"),
			output: "aaa",
		},
		{
			ID:     NewID(),
			output: "",
		},
		{
			ID:     ID("bbb"),
			output: "bbb",
		},
	}

	for _, tc := range testcases {
		if tc.ID.String() != tc.output {
			t.Error("入力と異なります", tc.ID.String(), tc.output)
		}
	}

}

func TestID_IsNew(t *testing.T) {
	testcases := []struct {
		ID    model.ID
		isNew bool
	}{
		{
			// テストIDの仕様による
			ID:    ID(""),
			isNew: true,
		},
		{
			// コレは本質的なテストではないかもしれない
			ID:    NewID(),
			isNew: true,
		},
		{
			ID:    ID("aaa"),
			isNew: false,
		},
	}

	for _, tc := range testcases {
		if tc.ID.IsNew() != tc.isNew {
			t.Error("入力と異なります", "get", tc.ID.IsNew(), "except", tc.isNew)
		}
	}
}

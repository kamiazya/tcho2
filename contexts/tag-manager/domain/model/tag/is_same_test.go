package tag

import (
	"testing"

	"bitbucket.org/kamiazya/tcho2/contexts/tag-manager/domain/values/stringid"
)

func TestIsSame(t *testing.T) {

	testcases := []struct {
		name   string
		a      Tag
		b      Tag
		except bool
	}{
		{
			name: "同じタグ",
			a: Tag{
				ID:   stringid.ID("a"),
				Name: "hoge",
			},
			b: Tag{
				ID:   stringid.ID("a"),
				Name: "hoge",
			},
			except: true,
		},
		{
			name: "同じタグだが名前が異なる",
			a: Tag{
				ID:   stringid.ID("a"),
				Name: "hoge",
			},
			b: Tag{
				ID:   stringid.ID("a"),
				Name: "fuga",
			},
			except: true,
		},
		{
			name: "IDが与えられていないタグ1",
			a: Tag{
				ID:   stringid.ID("a"),
				Name: "hoge",
			},
			b: Tag{
				ID:   stringid.NewID(),
				Name: "fuga",
			},
			except: false,
		},
		{
			name: "IDが与えられていないタグ2",
			a: Tag{
				ID:   stringid.NewID(),
				Name: "hoge",
			},
			b: Tag{
				ID:   stringid.NewID(),
				Name: "fuga",
			},
			except: false,
		},
		{
			name: "IDが与えられていないタグ3",
			a: Tag{
				ID:   stringid.NewID(),
				Name: "hoge",
			},
			b: Tag{
				ID:   stringid.ID("b"),
				Name: "fuga",
			},
			except: false,
		},
	}

	for _, tc := range testcases {
		if IsSame(tc.a, tc.b) != tc.except {
			t.Error(tc.name)
		}
	}
}

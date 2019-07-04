package main

import (
	"context"

	"github.com/k0kubun/pp"

	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/model/tag"
	"github.com/kamiazya/tcho2/contexts/tag-manager/infrastructure/database/inmem"
)

func main() {
	repo, err := inmem.NewTagsRepository(
		inmem.TagsRepoInitialCap(20),
		inmem.TagsRepoWithTags([]tag.Tag{
			{Name: "InitialTag1"},
			{Name: "InitialTag2"},
		}),
	)
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()

	repo.Store(ctx, &tag.Tag{
		Name: "TagA",
	})

	BID, _ := repo.Store(ctx, &tag.Tag{Name: "TagB"})

	repo.Update(ctx, BID, &tag.Tag{Name: "NewTagB"})

	tags, _ := repo.Find(ctx)

	pp.Println(tags)
}

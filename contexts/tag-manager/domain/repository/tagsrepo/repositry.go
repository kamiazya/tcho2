package tagsrepo

import (
	"context"

	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/model"
	"github.com/kamiazya/tcho2/contexts/tag-manager/domain/model/tag"
)

// Repository is the interface that governs the persistence of Tag.
// If a database dependent error occurs, it returns a DatabaseError error.
//
//
// Repository は、Tagの永続化を司るinterfaceである。
// データベース依存のエラーが発生した場合、DatabaseErrorのエラーを返す。
type Repository interface {
	// Store persists a new Tag and returns an ID.
	//
	// If Tag already has an ID, it returns ErrAleadyHasID.
	// If the name of Tag is not set, ErrTagNameIsRequired is returned.
	//
	//
	// Storeは、新規のTagを永続化し、IDを返す。
	//
	// もしTagが既にIDを持っていた場合、ErrAleadyHasIDを返す。
	// Tagの名前が未設定の場合、ErrTagNameIsRequiredを返す。
	Store(ctx context.Context, t *tag.Tag) (model.ID, error)

	// Update updates the Tag given in the argument.
	//
	// If the ID is newly issued, an error of ErrNotHaveID is returned.
	// If the tag with the corresponding ID can not be found, ErrNotFound error is returned.
	//
	//
	// Update は、引数で与えられたTagを更新する。
	//
	// IDが新規発行されたものの場合、ErrNotHaveIDのエラーを返す。
	// 該当するIDのタグが見つからなかった場合、ErrNotFoundのエラーを返す。
	Update(ctx context.Context, ID model.ID, t *tag.Tag) error

	// Find returns a slice of Tag that matches the given search criteria.
	// If an incorrect search condition is given, an error of ErrInvalidSearchCondition is returned.
	//
	//
	// Findは、与えられた検索条件に合致するTagのスライスを返す。
	//
	// もし不正な検索条件を与えた場合は、ErrInvalidSearchCondition のエラーを返す。
	Find(ctx context.Context, opts ...Option) (tags []tag.Tag, err error)
}

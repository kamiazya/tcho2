package repository

import "errors"

var (
	// ErrOptionAlreadySeted is an error that occurs when the same search condition is set more than once.
	//
	// ErrOptionAlreadySeted は、同じ検索条件を複数回セットしたときに発生するエラー。
	ErrOptionAlreadySeted = errors.New("ErrOptionAlreadySeted")

	// ErrInvalidSearchCondition is an error when an illegal value is set in the search condition.
	//
	// ErrInvalidSearchCondition は検索条件に不正な値がセットされたときのエラー。
	ErrInvalidSearchCondition = errors.New("ErrInvalidSearchCondition")

	// ErrAleadyHasID is an error that occurs when a model with an ID is newly saved.
	//
	// ErrAleadyHasID は、IDを持ったモデルが新規で保存されたときに発生するエラー。
	ErrAleadyHasID = errors.New("ErrAleadyHasID")

	// ErrNotHaveID is an error that occurs when processing fails because the model does not have an ID.
	//
	// ErrNotHaveID は、モデルがIDを持っておらず処理が失敗したときに起こるエラー。
	ErrNotHaveID = errors.New("ErrNotHaveID")

	// ErrNotFound is an error that occurs when the corresponding model can not be found.
	//
	// ErrNotFound は、該当のモデルが見つからなかったときに起こるエラー。
	ErrNotFound = errors.New("ErrNotFound")
)

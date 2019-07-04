package tagsrepo

import "errors"

var (
	// ErrTagNameIsRequired is an error that occurs when the value of Name is empty when saving Tag.
	//
	// ErrTagNameIsRequired は、 Tag を保存するときに、 Name の値が空だったときに発生するエラー。
	ErrTagNameIsRequired = errors.New("ErrTagNameIsRequired")
)

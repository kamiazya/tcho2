package model

import "fmt"

// ID is an interface representing a unique identifier of the entity.
//
// IDは、エンティティのユニークな識別子を表すinterfaceである。
type ID interface {
	IsNew() bool
	fmt.Stringer
}

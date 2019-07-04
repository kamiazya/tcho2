package repository

// DatabaseError is a database layer dependent error.
//
// DatabaseErrorはデータベース層に依存したエラー。
type DatabaseError struct {
	error
}

// Cause returns the contents of the actual error.
//
// Cause は、実際のエラーの内容を返す。
func (err DatabaseError) Cause() error {
	return err.error
}

package tag

// IsSame returns the value judging whether the tag given by the argument is the same or not.
//
// 引数で与えられたタグが同じものかどうかを判定した値を返す。
// もしnilが与えられた場合 falseを返す。
func IsSame(a, b Tag) bool {
	if a.IsNew() || b.IsNew() {
		return false
	}
	return a.String() == b.String()
}

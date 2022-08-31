package isFn

func IsTrue(v []bool) bool {
	if len(v) == 0 {
		return false
	}
	return v[0]
}

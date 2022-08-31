package isFn

func IsEmptyMap(v map[any]any) bool {
	if v == nil || len(v) == 0 {
		return true
	} else {
		return false
	}
}

func HasValueMap(v map[any]any) bool {
	if v == nil || len(v) == 0 {
		return false
	} else {
		return true
	}
}

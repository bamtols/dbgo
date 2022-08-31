package cliFn

func NilDefault[T any](v *T, def T) *T {
	if v == nil {
		return &def
	} else {
		return v
	}
}

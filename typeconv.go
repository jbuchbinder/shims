package shims

// ArrayConvert converts an array of one type to an array of another via an
// arbitrary transform function.
func ArrayConvert[F any, T any](in []F, f func(F) T) []T {
	out := make([]T, 0)
	for _, v := range in {
		out = append(out, f(v))
	}
	return out
}

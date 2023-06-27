package shims

// IfElse is a convenience function for the missing trinary operator.
func IfElse[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

package shims

// SingleValueDiscardError is used to wrap a single value function and
// discard the additional error.
func SingleValueDiscardError[T any](ret T, err error) T {
	return ret
}

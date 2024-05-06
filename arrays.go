package shims

// InArray returns a bool value indicating whether value is found in corpus.
func InArray[T comparable](value T, corpus []T) bool {
	for _, v := range corpus {
		if v == value {
			return true
		}
	}
	return false
}

// MatcherArray returns all members that match using an arbitrary matching
// function.
func MatcherArray[T any](in []T, f func(check T) bool) []T {
	o := make([]T, 0)
	for _, v := range in {
		if f(v) {
			o = append(o, v)
		}
	}
	return o
}

// MatcherArrayOne returns the first array member that matches using an
// arbitrary function. It returns a second parameter indicating whether or
// not a match had been made, to differentiate from the zero value for the
// type.
func MatcherArrayOne[T any](in []T, f func(check T) bool) (T, bool) {
	for _, v := range in {
		if f(v) {
			return v, true
		}
	}
	return *new(T), false
}

// ObjectArrayFunc executes an arbitrary function to reduce an array of
// objects to an array of arbitrary values.
func ObjectArrayFunc[T any, O any](in []T, f func(reduce T) O) []O {
	o := make([]O, 0)
	for _, v := range in {
		o = append(o, f(v))
	}
	return o
}

// RemoveArrayIndex removes an element from an array
func RemoveArrayIndex[T any](s []T, index int) []T {
	if index == len(s)-1 {
		return s[:index]
	}
	if len(s) == 1 {
		return make([]T, 0)
	}
	return append(s[:index], s[index+1:]...)
}

// FindArrayMember locates an array element by arbitrary criteria
func FindArrayMember[T any](in []T, f func(check T) bool) (int, bool) {
	for k, v := range in {
		if f(v) {
			return k, true
		}
	}
	return 0, false
}

// ArrayOfElements creates an array from a type.
func ArrayOfElements[T any](el T) []T {
	return make([]T, 0)
}

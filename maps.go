package shims

import "fmt"

// MapCopy creates a copy of a map
func MapCopy[T any, K comparable](in map[K]T) map[K]T {
	out := make(map[K]T, 0)
	for k := range in {
		out[k] = in[k]
	}
	return out
}

// Keys returns an array of keys in an arbitrary map.
func Keys[T any, K comparable](in map[K]T) []K {
	out := make([]K, 0)
	for k := range in {
		out = append(out, k)
	}
	return out
}

// Values returns an array of values in an arbitrary keyed map.
func Values[T any, U comparable](in map[U]T) []T {
	out := make([]T, 0)
	for _, v := range in {
		out = append(out, v)
	}
	return out
}

// MatcherKey returns a list of matched keys from an arbitrarily keyed map
func MatcherKey[T any, K comparable](in map[K]T, f func(check T) bool) []K {
	out := []K{}
	for k, o := range in {
		if f(o) {
			out = append(out, k)
		}
	}
	return out
}

// MatcherKeyOne returns the first matched key from an arbitrarily keyed map
func MatcherKeyOne[T any, K comparable](in map[K]T, f func(check T) bool) K {
	for k, o := range in {
		if f(o) {
			return k
		}
	}
	return *new(K)
}

// MatcherValue returns the matched values from an arbitrarily keyed map
func MatcherValue[T any, K comparable](in map[K]T, f func(check T) bool) ([]T, error) {
	out := make([]T, 0)
	for _, o := range in {
		if f(o) {
			out = append(out, o)
		}
	}
	var err error
	if len(out) == 0 {
		err = fmt.Errorf("not found")
	}
	return out, err
}

// MatcherValue returns the first matched value from an arbitrarily keyed map
func MatcherValueOne[T any, K comparable](in map[K]T, f func(check T) bool) (T, error) {
	for _, o := range in {
		if f(o) {
			return o, nil
		}
	}
	return *new(T), fmt.Errorf("not found")
}

// MergeMap merges two of the same type of maps together, overwriting values
// with the additional data.
func MergeMap[V comparable, T any](orig map[V]T, addl map[V]T) map[V]T {
	out := make(map[V]T)
	for k, v := range orig {
		out[k] = v
	}
	for k, v := range addl {
		out[k] = v
	}
	return out
}

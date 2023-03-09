package shims

import "fmt"

// Keys returns an array of keys in an arbitrary map.
func Keys[T any, K comparable](in map[K]T) []K {
	out := make([]K, 0)
	for k, _ := range in {
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

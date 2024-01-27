package factory

import (
	"fmt"
	"sync"
)

type Factory[T any] struct {
	dummy      func() T
	mutex      *sync.Mutex
	factoryMap map[string]func() T
}

// New creates a factory which stores and instantiates objects of
// the specified interface type. A "dummy" version needs to be passed
// to avoid Go's error checking mechanism from being triggered.
func New[T any](dummy T) Factory[T] {
	return Factory[T]{
		dummy: func() T { return dummy },
		mutex: &sync.Mutex{},
	}
}

// Register registers a new object for instantiation.
func (fac *Factory[T]) Register(name string, f func() T) {
	fac.mutex.Lock()
	defer fac.mutex.Unlock()

	if len(fac.factoryMap) == 0 {
		fac.factoryMap = map[string]func() T{}
	}
	fac.factoryMap[name] = f
}

// Instantiate is a factory function which instantiates an instance
// of an object with autodiscovery.
func (fac *Factory[T]) Instantiate(name string) (T, error) {
	fac.mutex.Lock()
	defer fac.mutex.Unlock()

	f, ok := fac.factoryMap[name]
	if !ok {
		return fac.dummy(), fmt.Errorf("unable to instantiate %s instance", name)
	}
	return f(), nil
}

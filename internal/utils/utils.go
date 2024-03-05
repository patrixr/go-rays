package utils

import "sync"

type Range struct {
	Min float64
	Max float64
}

func CreateRange(min float64, max float64) Range {
	return Range{Min: min, Max: max}
}

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// ------------------------------------------------

type Protected[T any] struct {
	Value T
	mx    *sync.Mutex
}

func Protect[T any](value T) Protected[T] {
	return Protected[T]{
		Value: value,
		mx:    &sync.Mutex{},
	}
}

func (p *Protected[T]) Take() T {
	p.mx.Lock()
	return p.Value
}

func (p *Protected[T]) Release() {
	p.mx.Unlock()
}

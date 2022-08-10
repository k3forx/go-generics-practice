package type_parameter

import "golang.org/x/exp/constraints"

func Min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

type Tree[T interface{}] struct {
	left, right *Tree[T]
	value       T
}

func (t *Tree[T]) Init(x T) *Tree[T] {
	return &Tree[T]{value: x}
}

func (t *Tree[T]) GetLeftAndRight() []*Tree[T] {
	return []*Tree[T]{t.left, t.right}
}

var StringTree Tree[string]

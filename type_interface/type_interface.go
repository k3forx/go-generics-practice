package type_interface

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Scale[S ~[]E, E constraints.Integer](s S, c E) S {
	r := make([]E, len(s))
	for i, v := range s {
		s[i] = v * c
	}
	return r
}

type Point []int32

func (p Point) String() string {
	return "String() method is called"
}

func ScaleAndPoint(p Point) {
	r := Scale(p, 2)
	fmt.Println(r.String())
}

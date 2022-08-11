package main

import (
	"fmt"

	"github.com/k3forx/go-generics-practice/type_interface"
	"github.com/k3forx/go-generics-practice/type_parameter"
)

func main() {
	x := type_parameter.GMin(2, 3)
	fmt.Printf("x: %+v\n", x)

	fmin := type_parameter.GMin[float64]
	fmt.Printf("y: %+v\n", fmin(2.71, 3.14))

	stringTree := type_parameter.StringTree.Init("init value")
	fmt.Printf("stringTree: %+v\n", stringTree)

	p := []int32{1, 2, 3}
	type_interface.ScaleAndPoint(p)
}

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("type: %T Value: %v\n", z, z)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	x, y := 3, 4
	fmt.Printf("x is of type %T\n", x)
	f1 := 1.0
	fmt.Printf("f1 is of type %T\n", f1)
	f = math.Sqrt(float64(x*x + x*y))
	var z = uint(f)
	fmt.Println(x, y, z)
}

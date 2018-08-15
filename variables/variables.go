package main

import "fmt"

// Pi rate for circle
const Pi = 3.14

const (
	//Big a number whose binary number is 1 followed by 100 zeros
	Big = 1 << 100
	//Small Shift BIG right again 99 places, it will be 1<<1, or 2
	Small = Big >> 99
)

var c, python, java bool

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

var (
	a = 1
	b = 2
)

func main() {
	const World = "世界"
	var i int
	var p, q int = 1, 2
	k := 3
	s1, s2 := "hello", "world"
	fmt.Print(i, p, q, k, c, s1, s2, python, java, "\n")
	fmt.Println("Happy", Pi, "Day")
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

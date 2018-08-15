package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/cmplx"
)

func euler() {
	// fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	// fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4

	c := int(math.Round(math.Sqrt(float64(a*a + b*b))))
	fmt.Println(c)
}

func readFile() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}
	return result
}
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	// euler()
	// triangle()
	// readFile()
	// fmt.Println(eval(100, 8, "/"))
	// fmt.Println(grade(0), grade(50), grade(60), grade(101), grade(81))
}

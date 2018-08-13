package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))
	// fmt.Println("My favorite number is", rand.Intn(10))
	// fmt.Println("My favorite number is", rand.Intn(10))
	// fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("Give PI is ", math.Pi)

}

package main

import (
	"fmt"
)

func run() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	var s22 []int
	fmt.Println(s22, len(s22), cap(s22))
	if s22 == nil {
		fmt.Println("nil")
	}

}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[2:6]
	printSlice(s)
	printSlice(arr[:6])
	printSlice(arr[:])
	printSlice(arr[:6])

}
func updateSlice(s []int) {
	s[0] = 100
}
func printSlice(s []int) {
	fmt.Printf("P: Len = %d cap=%d %v\n", len(s), cap(s), s)
}

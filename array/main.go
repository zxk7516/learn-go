package main

import (
	"fmt"
)

func fun() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "xxx"
	fmt.Println(a, b)
	fmt.Println(names)

	// 切片文法： 没有长度的数组
	q := []int{2, 3, 5, 7, 11, 12}
	fmt.Println(q)

	ss := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(ss)

	sss := []int{2, 3, 5, 7, 11, 13}

	sss = sss[1:4]
	fmt.Println(sss)

	sss = sss[:2]
	fmt.Println(sss)

	sss = sss[1:]
	fmt.Println(sss)
}
func printArr(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1, arr2, arr3)
	var grid [4][5]int
	fmt.Println(grid)

	// for i := 0; i < len(arr3); i++ {
	// fmt.Println(arr3[i])
	// }
	printArr(arr3)
}

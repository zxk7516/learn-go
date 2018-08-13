package main

import (
	"fmt"
)

func main() {
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

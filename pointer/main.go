package main

import "fmt"

func run() {
	i, j := 42, 2701

	p := &i // point to i
	fmt.Printf("%v", p)
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	fmt.Println(mySlice)
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func passByVal(a int) {
	a++
}
func passByRef(a *int) {
	*a++
}
func swap(a, b *int) {
	*b, *a = *a, *b
}
func main() {
	var a int = 3
	passByVal(a)
	println(a)
	passByRef(&a)
	println(a)
	b := 9
	swap(&a, &b)

	// run()
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

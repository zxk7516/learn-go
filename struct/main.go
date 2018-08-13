package main

import (
	"fmt"
)

// Vertex 大
type Vertex struct {
	X int
	Y int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (p person) updateName(newFirstName string) person {
	p.firstName = newFirstName
	return p
}

func (p *person) updateName2(newName string) {
	(*p).firstName = newName
}

func (p *person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	p1 := person{"Martice", "Tseng", contactInfo{"abc@t.co", 49000}}
	p1.print()
	p1 = p1.updateName("Boybird")
	(&p1).updateName2("Boybird2")
	p1.print()

	p2 := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 2323,
		},
	}
	p2.print()
	p2.updateFirstName("Jimmy")
	p2.print()

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Printf("%+v \n", alex)

	// v := Vertex{1, 2}
	// fmt.Println(v.X)
	// p := &v
	// p.X = 1e9
	// fmt.Println(v)
}

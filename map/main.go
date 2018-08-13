package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "ff0000",
		"green": "4bf745",
	}

	colors["green"] = "00ff00"
	fmt.Println(colors)
	colors["blue"] = "0000ff"
	delete(colors, "green") // 删除 key
	fmt.Println(colors)
	printMap(colors)

	var codes map[string]string
	fmt.Println(codes)

	m := make(map[int]string)
	m[10] = "#ffffff"
	fmt.Println(m)
}

func printMap(c map[string]string) {
	println("\n--------printMap-------")
	for color, hex := range c {
		fmt.Println("hex code for ", color, " is ", hex)
	}
	println("--------printMap-------\n")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func test() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

}

// go 语言中的 while 也定作 for
func while() {
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)
}

// 无限循环
func forever() {
	for {
	}
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	if len(result) == 0 {
		return "0"
	}
	return result
}

func printfile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	// fmt.Println(
	// 	convertToBin(5),  // 101
	// 	convertToBin(13), // 1101
	// 	convertToBin(72387885),
	// 	convertToBin(0),
	// )
	printfile("/Users/mt/Code/go/src/github.com/zxk7516/examples/README.md")
}

package main

import (
	"fmt"
	"io"
	"math"
	"reflect"
	"runtime"
	"strings"
)

type MyReader struct{}

func (mr MyReader) Read(b []byte) (n int, err error) {
	astrs := []byte{65}
	copy(b, astrs)
	return 1, nil
}

func read() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with arg"+"(%d %d)\n", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for _, n := range numbers {
		s += n
	}
	return s
}
func main() {
	// mr := MyReader{}
	// io.Copy(os.Stdout, mr)

	fmt.Println(sum(1, 3, 5, 7, 9))

	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

}

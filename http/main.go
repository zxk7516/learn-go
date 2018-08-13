package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// resp, err := http.Get("http://sign.hxhuiyuan.com")
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	// for k, v := range resp.Header {
	// 	fmt.Printf("k=%v, v=%v", k, v)
	// }
	// fmt.Printf("resp status %s,statusCode %d\n", resp.Status, resp.StatusCode)
	// fmt.Printf("resp Proto %s\n", resp.Proto)
	// fmt.Printf("resp content length %d\n", resp.ContentLength)
	// fmt.Printf("resp transfer encoding %v\n", resp.TransferEncoding)
	// fmt.Printf("resp Uncompressed %t\n", resp.Uncompressed)
	// fmt.Println(reflect.TypeOf(resp.Body)) // *http.gzipReader

	// buf := bytes.NewBuffer(make([]byte, 0, 512))
	// buf.ReadFrom(resp.Body)
	// fmt.Println(string(buf.Bytes()))

	// buf := []byte{}
	// buf := make([]byte, 99999)
	// resp.Body.Read(buf)
	// fmt.Println(string(buf))

	// buf2 := bytes.NewBuffer(make([]byte, 0, 512))
	// buf2.ReadFrom(resp.Body)
	// fmt.Println(string(buf2.Bytes()))

	lw := logWriter{}
	io.Copy(lw, resp.Body)

	// io.Copy(os.Stdout, resp.Body)
}

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {

	f, err := os.OpenFile("html.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return 0, err
	}

	n, err := f.Write(bs)
	fmt.Println(len(bs), n, len(bs) < n)

	if err == nil && n < len(bs) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return len(bs), err
}

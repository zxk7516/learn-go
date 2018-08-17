package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/zxk7516/examples/crawler/engine"
	"github.com/zxk7516/examples/crawler/zhenai/parser"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"regexp"
)

func getHtmlContents(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		panic("status error")
	}
	e := determineEncoding(resp.Body)
	// utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	return all
}
func main() {
	//url := "http://www.zhenai.com/zhenghun"
	//zhenhun := getHtmlContents(url)
	//printCityList(zhenhun)
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})

}

func determineEncoding(
	r io.Reader,
) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//for _,subMatch := range m{
		//	fmt.Printf("%s ",subMatch)
		//}
		fmt.Printf("%s -> %s\n", m[2], m[1])
	}
	fmt.Printf("Matched found: %d\n", len(matches))
}

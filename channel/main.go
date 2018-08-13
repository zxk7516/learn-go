package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.qq.com",
		"http://www.163.com",
		"http://www.baidu.com",
		"http://lol.duowan.com/",
		"http://www.huya.net",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down! error:", err)
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}

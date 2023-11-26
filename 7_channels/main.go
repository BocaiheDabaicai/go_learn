package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.baidu.com",
		"https://www.weibo.com",
		"https://www.zhihu.com",
		"https://www.youdao.com",
		"https://www.iqiyi.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
		//fmt.Println(<-c)
	}

	for {
		go checkLink(<-c, c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

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

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}

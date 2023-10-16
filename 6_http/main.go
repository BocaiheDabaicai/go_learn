package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	/*	// 1. 直接读取请求数据
		bs := make([]byte, 99999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/

	/*	// 2. 使用输入输出对象复制Body数据，由os对象进行打印
		io.Copy(os.Stdout, resp.Body)
	*/
	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

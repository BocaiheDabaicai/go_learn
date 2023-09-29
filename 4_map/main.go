package main

import (
	"fmt"
)

func main() {
	// 定义
	//var colors map[string]string

	// 定义二
	//colors := make(map[string]string)

	// 赋值定义
	/*colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}*/

	//colors["while"] = "#ffffff"
	//colors[strconv.Itoa(10)] = "#ffffff"
	//delete(colors, strconv.Itoa(10))

	//fmt.Println(colors)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

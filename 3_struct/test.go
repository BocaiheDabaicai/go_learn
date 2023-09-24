package main

import "fmt"

func main() {
	mySlice := []string{"Hello", "my", "dear", "friend"}

	updateSlice(mySlice)

	fmt.Println(mySlice)

}

func updateSlice(s []string) {
	s[0] = "Good Morning"
}

package main

import "fmt"

func main() {
	a := 1
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
}

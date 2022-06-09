package main

import "fmt"

func main() {
	a := 1
	fmt.Println(&a)
	var p = &a
	fmt.Println(*p)
}

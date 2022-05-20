package main

func foo(m0 int) *int {
	var m1 int = 11
	return &m1
}

func main() {
	m := foo(100)
	println(*m) //11
}

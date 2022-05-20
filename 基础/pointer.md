package main

import "fmt"

// pointer 指针
// [Go语言基础之指针](https://www.liwenzhou.com/posts/Go/07_pointer/)

func main() {
	p := 520
	o := &p
	i := &o
	n := &i
	t := *o
	e := *i
	fmt.Println(o)
	fmt.Println(i)
	fmt.Println(n)
	fmt.Printf("o:%p type:%T\n", o, o)
	fmt.Println(t)
	fmt.Println(e)

	//NewAndMake()
	//执行上面的代码会引发panic，为什么呢？
	//在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
	//而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
	//要分配内存，就引出来今天的new和make。
	//Go语言中new和make是内建的两个函数，主要用来分配内存。

	UseNew()

	UseMake()
}

// NewAndMake new and make
//new与make的区别
//二者都是用来做内存分配的。
//make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
//而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
func NewAndMake() {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["jokereven"] = 100
	fmt.Println(b)
}

func UseNew() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}

func UseMake() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["jokereven"] = 100
	fmt.Println(b)
}

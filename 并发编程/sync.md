```Go
package main

// https://www.topgoer.com/并发编程/sync.html

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("Hello Goroutine!")
}
func main() {
	wg.Add(1)
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	wg.Wait()
}

```

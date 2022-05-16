package main

import "fmt"

// [channel](https://www.topgoer.com/并发编程/channel.html)
// channel
// var ch chan int

// create a channel
func createChannel() {
	var ch chan int
	fmt.Println(ch)
}

// initialize a channel
func initializeChannel() {
	ch := make(chan int)
	fmt.Println(ch)
}

// 关于关闭通道需要注意的事情是，只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。
// 通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。
// 关闭后的通道有以下特点：
//1.对一个关闭的通道再发送值就会导致panic。
//2.对一个关闭的通道进行接收会一直获取值直到通道为空。
//3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
//4.关闭一个已经关闭的通道会导致panic。

// send a message to a channel
func sendMessage() {
	// 发送 接收 与 关闭
	ch := make(chan int) // deadlock
	// send
	ch <- 10
	// get the message
	x := <-ch // 从ch中接收值并赋值给变量x
	<-ch      // 从ch中接收值，忽略结果
	fmt.Println(x)
	// close
	close(ch)
}

// deadlock

//为什么会出现deadlock错误呢？
//因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值。
//就像你住的小区没有快递柜和代收点，快递员给你打电话必须要把这个物品送到你的手中，简单来说就是无缓冲的通道必须有接收才能发送。
//上面的代码会阻塞在ch <- 10这一行代码形成死锁，那如何解决这个问题呢？

// 通过 goroutine 解决 deadlock
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func ByGoroutineResolvedDeadlock() {
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}

// BufferChannels 有缓冲的通道
func BufferChannels() {
	ch := make(chan int, 1)
	ch <- 10
	fmt.Println("by BufferChannels send successfully")
}

// CloseChannel 关闭通道
//可以通过内置的close()函数关闭channel（如果你的管道不往里存值或者取值的时候一定记得关闭管道）
func CloseChannel() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("CloseChannel over")
}

// WhenChannelClosed 如何优雅的从通道循环取值
//当通过通道发送有限的数据时，我们可以通过close函数关闭通道来告知从该通道接收值的goroutine停止等待。
//当通道被关闭时，往该通道发送值会引发panic，从该通道里接收的值一直都是类型零值。那如何判断一个通道是否被关闭了呢？
func WhenChannelClosed() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}

// 单向通道
//有的时候我们会将通道作为参数在多个任务函数间传递，
//很多时候我们在不同的任务函数中使用通道都会对其进行限制，
//比如限制通道在函数中只能发送或只能接收。

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func dire() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

func main() {
	createChannel()     //<nil>
	initializeChannel() //0x11814340
	//sendMessage()       // deadlock
	ByGoroutineResolvedDeadlock()
	BufferChannels()
	CloseChannel()
	WhenChannelClosed()
	dire()
}

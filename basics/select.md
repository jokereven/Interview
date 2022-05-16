package main

import "fmt"

// [Go语言在select语句中实现优先级](https://www.liwenzhou.com/posts/Go/priority_in_go_select)
// select

// Basics
//Go 语言中的 select 关键字也能够让当前 goroutine 同时等待ch1 的可读和ch2的可写，
//在ch1和ch2状态改变之前，select 会一直阻塞下去，直到其中的一个 channel 转为就绪状态时执行对应case分支的代码。
//如果多个channel同时就绪的话则随机选择一个case执行。
func Basics() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	select {
	case <-ch1:
		fmt.Println("github.com")
	case ch2 <- 1:
		fmt.Println("jokereven")
	}
}

func EmptySelect() {
	// 空 select 无 case
	//空的 select 语句会直接阻塞当前的goroutine，使得该goroutine进入无法被唤醒的永久休眠状态。
	select {}
}

func OnlyOneSelect() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	select {
	// 可读
	case <-ch1:
		fmt.Println("Only one select")
	}
}

//总结
//select 不存在任何的 case：永久阻塞当前 goroutine
//select 只存在一个 case：阻塞的发送/接收
//select 存在多个 case：随机选择一个满足条件的case执行
//select 存在 default，其他case都不满足时：执行default语句中的代码

func SelectUseDefault() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	select {
	case <-ch1:
		fmt.Println("select default")
	default:
		fmt.Println("this is select default")
	}
}

//已知，当select 存在多个 case时会随机选择一个满足条件的case执行。
//现在我们有一个需求：我们有一个函数会持续不间断地从ch1和ch2中分别接收任务1和任务2，
//如何确保当ch1和ch2同时达到就绪状态时，优先执行任务1，在没有任务1的时候再去执行任务2呢？
// worker01
func worker01(ch1, ch2 <-chan int, stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		case job1 := <-ch1:
			fmt.Println(job1)
		default:
			select {
			case job2 := <-ch2:
				fmt.Println(job2)
			default:
			}
		}
	}
}

//[Go 语言中的label使用](https://juejin.cn/post/7030996392487157773)
//[Go语言基础之流程控制](https://www.liwenzhou.com/posts/Go/04_basic)
// worker02 使用label
func worker02(ch1, ch2 <-chan int, stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		case job1 := <-ch1:
			fmt.Println(job1)
		case job2 := <-ch2:
		priority:
			for {
				select {
				case job1 := <-ch1:
					fmt.Println(job1)
				default:
					break priority
				}
			}
			fmt.Println(job2)
		}
	}
}

func main() {
	Basics() //jokereven
	OnlyOneSelect()
	SelectUseDefault()
}

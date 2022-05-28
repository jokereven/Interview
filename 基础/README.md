## Golang 基础

### 1.1.1. select 是随机的还是顺序的?
答: select 会随机选择一个可用通道做收发操作

matlab: [select](./select.md)
[channel](./channel.md)

### 1.1.2. Go 语言局部变量分配在栈还是堆？
答: Go 语言编译器会自动决定把一个变量放在栈还是放在堆，编译器会做逃逸分析，当发现变量的作用域没有跑出函数范围，就可以在栈上，反之则必须分配在堆。

参考: [go 语言局部变量分配在栈还是堆](https://www.jianshu.com/p/4e3478e9d252)
[局部变量分配在栈还是堆](./local_variable_heap_or_stack.md)

matlab: [指针](./pointer.md)

### 1.1.3. 简述一下你对 Go 垃圾回收机制的理解?
答:

v1.1 STW
v1.3 Mark STW, Sweep 并行
v1.5 三色标记法
v1.8 hybrid write barrier(混合写屏障：优化 STW)

bilibili: [golang 垃圾回收-原理篇](https://www.bilibili.com/video/BV1NL411A7wX)
[25.三色标记算法原理解决漏标问题](https://www.bilibili.com/video/BV1Ar4y1b7Rd)

参考: [Golang 垃圾回收剖析](http://legendtkl.com/2017/04/28/golang-gc/)
[[典藏版]Golang 三色标记、混合写屏障 GC 模式图文全分析](https://segmentfault.com/a/1190000022030353)

### 1.1.4. 简述一下 golang 的协程调度原理?
答:

M(machine): 代表着真正的执行计算资源，可以认为它就是 os thread（系统线程）。
P(processor): 表示逻辑 processor，是线程 M 的执行的上下文。
G(goroutine): 调度系统的最基本单位 goroutine，存储了 goroutine 的执行 stack 信息、goroutine 状态以及 goroutine 的任务函数等。

参考: [Goroutine](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/)

bilibili: [【os 浅尝】话说进程和线程~](https://www.bilibili.com/video/BV1H541187UH)
[【Go 面试】Go goroutine 和线程的区别？](https://www.bilibili.com/video/BV1Kr4y1W72a)

### 1.1.5. 介绍下 golang 的 runtime 机制?
答: Runtime 负责管理任务调度，垃圾收集及运行环境。同时，Go 提供了一些高级的功能，如 goroutine, channel, 以及 Garbage collection。这些高级功能需要一个 runtime 的支持. runtime 和用户编译后的代码被 linker 静态链接起来，形成一个可执行文件。这个文件从操作系统角度来说是一个 user space 的独立的可执行文件。 从运行的角度来说，这个文件由 2 部分组成，一部分是用户的代码，另一部分就是 runtime。runtime 通过接口函数调用来管理 goroutine, channel 及其他一些高级的功能。从用户代码发起的调用操作系统 API 的调用都会被 runtime 拦截并处理。

Go runtime 的一个重要的组成部分是 goroutine scheduler。他负责追踪，调度每个 goroutine 运行，实际上是从应用程序的 process 所属的 thread pool 中分配一个 thread 来执行这个 goroutine。因此，和 java 虚拟机中的 Java thread 和 OS thread 映射概念类似，每个 goroutine 只有分配到一个 OS thread 才能运行。

参考: [go runtime 的机制如何](https://blog.csdn.net/xclyfe/article/details/50562349)

bilibili: [【Golang】信号量 - runtime 提供的等待队列](https://www.bilibili.com/video/BV1ZQ4y1f7go)

### 1.1.6. 如何获取 go 程序运行时的协程数量, gc 时间, 对象数, 堆栈信息?
答:调用接口 runtime.ReadMemStats 可以获取以上所有信息, 注意: 调用此接口会触发 STW(Stop The World)

参考: https://golang.org/pkg/runtime/#ReadMemStats

如果需要打入到日志系统, 可以使用 go 封装好的包, 输出 json 格式. 参考:

https://golang.org/pkg/expvar/
http://blog.studygolang.com/2017/06/expvar-in-action/
更深入的用法就是将得到的运行时数据导入到 ES 内部, 然后使用 Kibana 做 golang 的运行时监控, 可以实时获取到运行的信息(堆栈, 对象数, gc 时间, goroutine, 总内存使用等等), 具体信息可以看 ReadMemStats 的那个结构体

### 1.1.7. 介绍下你平时都是怎么调试 golang 的 bug 以及性能问题的?
答:

panic 调用栈
pprof
火焰图(配合压测)
使用 go run -race 或者 go build -race 来进行竞争检测
查看系统 磁盘 IO/网络 IO/内存占用/CPU 占用(配合压测)

### 1.1.8. 简单介绍下 golang 中 make 和 new 的区别
答: new(T) 是为一个 T 类型的新值分配空间, 并将此空间初始化为 T 的零值, 并返回这块内存空间的地址, 也就是 T 类型的指针 T, 该指针指向 T 类型值占用的那块内存. make(T) 返回的是初始化之后的 T, 且只能用于 slice, map, channel 三种类型. make(T, args) 返回初始化之后 T 类型的值, 且此新值并不是 T 类型的零值, 也不是 T 类型的指针 T, 而是 T 类型值经过初始化之后的引用.

参考1: [Go中make和new的区别](https://www.cnblogs.com/ghj1976/archive/2013/02/12/2910384.html)

参考2: [Go中make()和new()的区别](https://studygolang.com/articles/3496)

### 1.1.9. 简单说下Golang逃逸分析
参考: [Golang内存分配逃逸分析](https://www.iphpt.com/detail/137)

### 1.1.10. 无缓冲 Chan 的发送和接收是否同步?
答：

channel无缓冲时，发送阻塞直到数据被接收，接收阻塞直到读到数据。
channel有缓冲时，当缓冲满时发送阻塞，当缓冲空时接收阻塞。


### 1.1.11. Golang通过哪几种方式来实现并发控制,如何优雅的退出goroutine?
答：

chan 通过无缓冲通道来实现多 goroutine 并发控制
通过sync包中的WaitGroup 实现并发控制

退出：

使用for-range退出

使用,ok退出
使用退出通道退出
参考：

https://golangnote.com/topic/184.html

http://lessisbetter.site/2018/12/02/golang-exit-goroutine-in-3-ways/

### 1.1.12. Golang的interface的特性和技巧,举例一些优雅的实现?

答：

- 空接口（empty interface）

    - 空接口比较特殊，他不包含任何方法,但是他又可以表示任何类型

    - golang的所有基础类都实现了空接口，所有我们可以用[]interface表示结构不同的数组,比如:

    ```
    func main() {
            data := make([]interface{}, 3)
            intData := 1
            stringData := "abc"
            boolData := true
            data[0] = intData
            data[1] = stringData
            data[2] = boolData
            for _, v := range data {
                fmt.Println(v)
            }
        }
    ```


- 接口嵌套接口

    - 一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。

    - 比如接口 File 包含了 ReadWrite 和 Lock 的所有方法，它还额外有一个 Close() 方法。


    ```
        type ReadWrite interface {
            Read(b Buffer) bool
            Write(b Buffer) bool
        }
        type Lock interface {
            Lock()
            Unlock()
        }
        type File interface {
            ReadWrite
            Lock
            Close()
        }
    ```

- 类型的选择与断言

    - 一个接口类型的变量 varI 中可以包含任何类型的值，必须有一种方式来检测它的 动态 类型，即运行时在变量中存储的值的实际类型。在执行过程中动态类型可能会有所不同，但是它总是可以分配给接口变量本身的类型。通常我们可以使用类型断言 来测试在某个时刻 接口varI 是否包含类型 T 的值：

    ```
     v := varI.(T)
    ```

    - 类型断言可能是无效的，虽然编译器会尽力检查转换是否有效，但是它不可能预见所有的可能性。如果转换在程序运行时失败会导致错误发生。更安全的方式是使用以下形式来进行类型断言：

    ```
        if v, ok := varI.(T); ok {  // checked type
        assertion
            Process(v)
            return
        }
    ```

    > 参考： [https://blog.csdn.net/DB\_water/article/details/79068271](https://blog.csdn.net/DB_water/article/details/79068271)

### 1.1.13. Golang的方法集?

答：

> - 类型 T 方法集包含全部 receiver T 方法。
> - 类型 _T 方法集包含全部 receiver T +_ T 方法。
> - 如类型 S 包含匿名字段 T，则 S 和 \*S 方法集包含 T 方法。
> - 如类型 S 包含匿名字段 _T，则 S 和_ S 方法集包含 T + \*T 方法。
> - 不管嵌入 T 或 _T，_S 方法集总是包含 T + \*T 方法。

用实例 value 和 pointer 调用方法 (含匿名字段) 不受方法集约束，编译器总是查找全部方法，并自动转换 receiver 实参。

Go 语言中内部类型方法集提升的规则：

- 类型 T 方法集包含全部 receiver T 方法。

    ```
         package main

         import (
         "fmt"
         )

         type T struct {
         int
         }

         func (t T) test() {
         fmt.Println("类型 T 方法集包含全部 receiver T 方法。")
         }

         func main() {
         t1 := T{1}
         fmt.Printf("t1 is : %v\n", t1)
         t1.test()
         }
    ```

    输出结果：

    > t1 is : {1}
    >
    > 类型 T 方法集包含全部 receiver T 方法。

- 类型 _T 方法集包含全部 receiver T +_ T 方法。

    ```
    package main

    import (
        "fmt"
    )

    type T struct {
        int
    }

    func (t T) testT() {
        fmt.Println("类型 *T 方法集包含全部 receiver T 方法。")
    }

    func (t *T) testP() {
        fmt.Println("类型 *T 方法集包含全部 receiver *T 方法。")
    }

    func main() {
        t1 := T{1}
        t2 := &t1
        fmt.Printf("t2 is : %v\n", t2)
        t2.testT()
        t2.testP()
    }
    ```

    输出结果：

    > t2 is : &{1}
    >
    > 类型 \*T 方法集包含全部 receiver T 方法。
    >
    > 类型 _T 方法集包含全部 receiver_ T 方法。

    给定一个结构体类型 S 和一个命名为 T 的类型，方法提升像下面规定的这样被包含在结构体方法集中：

- 如类型 S 包含匿名字段 T，则 S 和 \*S 方法集包含 T 方法。

    这条规则说的是当我们嵌入一个类型，嵌入类型的接受者为值类型的方法将被提升，可以被外部类型的值和指针调用。

    ```
       package main

       import (
         "fmt"
       )

       type S struct {
         T
       }

       type T struct {
         int
       }

       func (t T) testT() {
         fmt.Println("如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。")
       }

       func main() {
         s1 := S{T{1}}
         s2 := &s1
         fmt.Printf("s1 is : %v\n", s1)
         s1.testT()
         fmt.Printf("s2 is : %v\n", s2)
         s2.testT()
       }
    ```

    输出结果：

    > s1 is : 1
    >
    > 如类型 S 包含匿名字段 T，则 S 和 \*S 方法集包含 T 方法。
    >
    > s2 is : &1
    >
    > 如类型 S 包含匿名字段 T，则 S 和 \*S 方法集包含 T 方法。

- 如类型 S 包含匿名字段 _T，则 S 和_ S 方法集包含 T + \*T 方法。

    这条规则说的是当我们嵌入一个类型的指针，嵌入类型的接受者为值类型或指针类型的方法将被提升，可以被外部类型的值或者指针调用。

    ```
     package main

     import (
         "fmt"
     )

     type S struct {
         T
     }

     type T struct {
         int
     }

     func (t T) testT() {
         fmt.Println("如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T 方法")
     }
     func (t *T) testP() {
         fmt.Println("如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 *T 方法")
     }

     func main() {
         s1 := S{T{1}}
         s2 := &s1
         fmt.Printf("s1 is : %v\n", s1)
         s1.testT()
         s1.testP()
         fmt.Printf("s2 is : %v\n", s2)
         s2.testT()
         s2.testP()
     }
    ```

    输出结果：

    > s1 is : 1
    >
    > 如类型 S 包含匿名字段 _T，则 S 和_ S 方法集包含 T 方法
    >
    > 如类型 S 包含匿名字段 _T，则 S 和_ S 方法集包含 \*T 方法
    >
    > s2 is : &1
    >
    > 如类型 S 包含匿名字段 _T，则 S 和_ S 方法集包含 T 方法
    >
    > 如类型 S 包含匿名字段 _T，则 S 和_ S 方法集包含 \*T 方法


> [叶落山城秋](https://www.iphpt.com/): 上面的其实很好理解!但是我懵逼的是什么呢,感觉上面 不管是 T 还是 _T 还是 S或者_ S 发现都可以!
>
> 那什么情况下是不可以呢 如果 是 [代码考题第10题](http://interview.wzcu.com/Golang/%E4%BB%A3%E7%A0%81%E8%80%83%E9%A2%98.html#%E4%BB%A5%E4%B8%8B%E4%BB%A3%E7%A0%81%E8%83%BD%E7%BC%96%E8%AF%91%E8%BF%87%E5%8E%BB%E5%90%97%EF%BC%9F%E4%B8%BA%E4%BB%80%E4%B9%88%EF%BC%9F) 如果上面是接口
>
> 发现如果是 不ok的.. 但是 var peo = Student{} 又是ok的! 那 var peo People = Student{} 这样写的意义在于?? 先记着吧. 如果哪位大佬比较明白这块,跪请留言告知!十分感谢

### 1.1.14. Golang的GMP模型?

- M(Work Thread) -- 表示操作系统的线程,它是被操作系统管理的线程,与POSIX中的标准线程非常类似
- G(Goroutine) -- 表示Goroutine,每一个Goroutine都包含堆栈,指令指针和其他用于调度的重要信息
- P(Processor) -- 表示调度的上下文,它可以被看做一个运行于线程M上的本地调度器

    三者关系:

    - 每一个运行的M都必须绑定一个P,线程M创建后会去检查并执行G(goroutine)对象
    - 每一个P保存着一个协程G的队列
    - 除了每个P自身保存的G的队列外,调度器还拥有一个全局的G队列
    - M从队列中提取G,并执行
    - P的个数就是`GOMAXPROCS`(最大256),启动时固定的,一般不修改
    - M的个数和P的个数不一定一样多(会有休眠的M或P不绑定M) (最大10000)
    - P是用一个全局数组(255)来保存的,并且维护着一个全局的P空闲链表

### 1.1.15. 用信道实现主程等待协程2s,如果超过2s,主程直接结束(不用sleep)?

```
func main() {
    start := time.Now()
    wait := make(chan int,1)
    go func() {
        fmt.Println("做点东西")
        time.Sleep(1*time.Second)
        wait<-2
    }()
    fmt.Println("这里是主程序")
    select {
    case nums:= <-wait:
        fmt.Println(nums)
    case <-time.After(2*time.Second):
        fmt.Println("2秒后")
    }
    fmt.Println(time.Since(start))
}
```

### 1.1.16. Golang里的结构体是否能进行比较?

先说答案, 部分情况下是可以比较的

比如:

1. 结构体属性为一些常规类型比如 int,string 可以进行比较
2. 如果结构体属性是一些引用类型,比如 切片,map 等,不可进行比较

假设结构体属性都是常规的类型

1. 如果是同一个结构体,那肯定是可以比较的(申明一个结构体)

    ```
     func() {
         type Test struct{
             A string
             B int
         }
         a := Test{
             A:"1",
             B:1,
         }
         b := Test{
             A:"1",
             B:1,
         }
         fmt.Println(a==b)
    }
    ```

2. 如果是两个结构体,但是属性是一样的,且顺序都一样,是可以比较的, **如果顺序不一样,则不能比较**

    ```
     func() {
         a := struct {
             a int
             b string
         }{
             a: 1,
             b: "1",
         }

         b := struct {
             a int
             b string
         }{
             a: 1,
             b: "1",
         }
         fmt.Println(a==b)
     }
    ```


另外,如果要比较两个结构体,可以用 `reflect.DeepEqual()` 方法进行比较

加载中...

## TODO

### command line

- [x] initiate
- [ ] participate
- [ ] redeem
- [ ] refund

### doc
- [x] ReadMe
- [ ] ReadMe_CN

Golang 基础

1.1.1. select 是随机的还是顺序的?
答: select 会随机选择一个可用通道做收发操作

matlab: [select](./select.md)
[channel](./channel.md)

1.1.2. Go 语言局部变量分配在栈还是堆？
答: Go 语言编译器会自动决定把一个变量放在栈还是放在堆，编译器会做逃逸分析，当发现变量的作用域没有跑出函数范围，就可以在栈上，反之则必须分配在堆。

参考: [go 语言局部变量分配在栈还是堆](https://www.jianshu.com/p/4e3478e9d252)
[局部变量分配在栈还是堆](./local_variable_heap_or_stack.md)

matlab: [指针](./pointer.md)

1.1.3. 简述一下你对 Go 垃圾回收机制的理解?
答:

v1.1 STW
v1.3 Mark STW, Sweep 并行
v1.5 三色标记法
v1.8 hybrid write barrier(混合写屏障：优化 STW)

bilibili: [golang 垃圾回收-原理篇](https://www.bilibili.com/video/BV1NL411A7wX)
[25.三色标记算法原理解决漏标问题](https://www.bilibili.com/video/BV1Ar4y1b7Rd)

参考: [Golang 垃圾回收剖析](http://legendtkl.com/2017/04/28/golang-gc/)
[[典藏版]Golang 三色标记、混合写屏障 GC 模式图文全分析](https://segmentfault.com/a/1190000022030353)

1.1.4. 简述一下 golang 的协程调度原理?
答:

M(machine): 代表着真正的执行计算资源，可以认为它就是 os thread（系统线程）。
P(processor): 表示逻辑 processor，是线程 M 的执行的上下文。
G(goroutine): 调度系统的最基本单位 goroutine，存储了 goroutine 的执行 stack 信息、goroutine 状态以及 goroutine 的任务函数等。

参考: [Goroutine](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/)

bilibili: [【os 浅尝】话说进程和线程~](https://www.bilibili.com/video/BV1H541187UH)
[【Go 面试】Go goroutine 和线程的区别？](https://www.bilibili.com/video/BV1Kr4y1W72a)

1.1.5. 介绍下 golang 的 runtime 机制?
答: Runtime 负责管理任务调度，垃圾收集及运行环境。同时，Go 提供了一些高级的功能，如 goroutine, channel, 以及 Garbage collection。这些高级功能需要一个 runtime 的支持. runtime 和用户编译后的代码被 linker 静态链接起来，形成一个可执行文件。这个文件从操作系统角度来说是一个 user space 的独立的可执行文件。 从运行的角度来说，这个文件由 2 部分组成，一部分是用户的代码，另一部分就是 runtime。runtime 通过接口函数调用来管理 goroutine, channel 及其他一些高级的功能。从用户代码发起的调用操作系统 API 的调用都会被 runtime 拦截并处理。

Go runtime 的一个重要的组成部分是 goroutine scheduler。他负责追踪，调度每个 goroutine 运行，实际上是从应用程序的 process 所属的 thread pool 中分配一个 thread 来执行这个 goroutine。因此，和 java 虚拟机中的 Java thread 和 OS thread 映射概念类似，每个 goroutine 只有分配到一个 OS thread 才能运行。

参考: [go runtime 的机制如何](https://blog.csdn.net/xclyfe/article/details/50562349)

bilibili: [【Golang】信号量 - runtime 提供的等待队列](https://www.bilibili.com/video/BV1ZQ4y1f7go)

1.1.6. 如何获取 go 程序运行时的协程数量, gc 时间, 对象数, 堆栈信息?
答:调用接口 runtime.ReadMemStats 可以获取以上所有信息, 注意: 调用此接口会触发 STW(Stop The World)

参考: https://golang.org/pkg/runtime/#ReadMemStats

如果需要打入到日志系统, 可以使用 go 封装好的包, 输出 json 格式. 参考:

https://golang.org/pkg/expvar/
http://blog.studygolang.com/2017/06/expvar-in-action/
更深入的用法就是将得到的运行时数据导入到 ES 内部, 然后使用 Kibana 做 golang 的运行时监控, 可以实时获取到运行的信息(堆栈, 对象数, gc 时间, goroutine, 总内存使用等等), 具体信息可以看 ReadMemStats 的那个结构体

1.1.7. 介绍下你平时都是怎么调试 golang 的 bug 以及性能问题的?
答:

panic 调用栈
pprof
火焰图(配合压测)
使用 go run -race 或者 go build -race 来进行竞争检测
查看系统 磁盘 IO/网络 IO/内存占用/CPU 占用(配合压测)

加载中...

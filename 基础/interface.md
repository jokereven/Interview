package main

import (
	"fmt"
)

// [interface-liwenzhou](https://www.liwenzhou.com/posts/Go/12-interface/)

// 类型断言
// [Go 类型断言是什么？](https://juejin.cn/post/6962522545039867912)
// 关于类型断言，用人话来说就是对类型的判断。
// 变量.(类型)。例如：i.(int)
// 变量,bool = 变量.(类型)。例如：num,ok = i.(int)。ok表示判断类型是否成功的意思。

func main() {
	// num ...
	var num interface{} = 10
	// i := num.(int)
	// i := num.(string)
	// fmt.Println("获取num的值:", i)
	// fmt.Println("获取num的类型:", reflect.TypeOf(i))

	// str := num.(string)
	str := num.(int)
	fmt.Println(str)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

}

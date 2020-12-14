package Func

import "fmt"

func init() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(100, 10) // 通过变量调用匿名函数

	// 自执行函数: 匿名函数定义完成()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}

// SayHi 普通函数
func SayHi() {
	fmt.Println("hello")
}

// Resdata 有返回函数
func Resdata() int {
	return 10
}

// Parmdata参数函数
func Parmdata(name string, age int) {
	fmt.Printf("name:%v, age:%v \n", name, age)
}

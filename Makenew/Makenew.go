package Makenew

import "fmt"

func init() {
	// 定义一个变量-但是没有初始化(内存还没有分配， 所以不能赋值)
	var b map[string]int
	b = make(map[string]int, 10)
	b["name"] = 100
	fmt.Printf("%v \n", b)
}

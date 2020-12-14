package main

import (
	"fmt"
	FUNC "godemo/Func" // 函数
	_ "godemo/Func"    // 函数
	_ "godemo/Makenew" // 分配内存
	_ "godemo/This"    // 指针
)

func main() {
	FUNC.SayHi()
	fmt.Printf("%v \n", FUNC.Resdata())
	FUNC.Parmdata("张德帅", 20)
}

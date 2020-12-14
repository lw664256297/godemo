package Func

import "fmt"

// SayHi 普通函数
func SayHi(){
	fmt.Println("hello")
}

// Resdata 有返回函数
func Resdata() int  {
	return 10
}

// Parmdata参数函数
func Parmdata(name string, age int)   {
	fmt.Printf("name:%v, age:%v \n", name, age)
}
package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// go 语言变量有初始值
func variableZeroValue(){
	var a int
	var s string
	fmt.Printf("%d %q \n",a,s)
}

func variableInitalValue(){
	var a,b int = 3,4
	var s string = "abc"
	fmt.Println(a,b,s)
}

// 类型推断，这里定义变量的时候后面并没有跟类型
func variableTypeDeduction(){
	var a,b ,c,s = 3,4,true,"def"
	fmt.Println(a,b,c,s)
}

// :=的方式可以不用写var
func variableShorter(){
	a,b,c,s := 3,4,true,"def"
	b = 5 // 这是第二次赋值，上面一行是第一次 定义+赋值，这个时候是必须是：=。
	fmt.Println(a,b,c,s)
}

// 方法之外的变量定义, 这些变量的作用域不是全局的，而是包内部的变量
var m = 2
var s = "s"
var b = true
// 这种写法不行 c:="c"，函数外面的变量，不能用 :=
var (
	i = 1
	j = "a"
	k = true
)// 这种写法就不需要写3个var了。


func main() {
	fmt.Println("hello")
	variableZeroValue()
	variableInitalValue()
	variableTypeDeduction()
	variableShorter()

	euler()

	triangle()

	consts()

	enums()
}


// go 的内建变量类型
// bool , string
// (u)int ,(u)int8,(u)int16,(u)int32,(u)int64,uintptr   u ,就是有无符号
// byte ,rune    rune 是字符型（32位）相当于char
// float32 , float64 , complex64 , complex128   复数
func euler(){
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}


// 强制类型转换
func triangle(){
	var a,b int = 3,4
	var c int
	//c = math.Sqrt(a*a + b*b)
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}


// 常量定义
const name = "zhangsan"
func consts(){
	const filename = "abc.txt"
	const a,b = 3,4
	var c int
	c = int(math.Sqrt( a*a + b*b))
	fmt.Println(filename,name,c)
}

// 枚举类型
func enums(){
	/*const(
		cpp = 0
		java = 1
		python = 2
		golang = 3
	)*/
	const(
		cpp = iota
		java
		python
		golang
	)
	fmt.Println(cpp,java,python,golang)
	// b kb mb gb tb pb
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

// 变量类型写在变量名之后
// 编译器可以推断变量类型
// 没有char, 只有rune
// 原生支持复数类型
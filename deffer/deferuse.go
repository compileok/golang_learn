package main

import "fmt"

func main() {
	fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println(4)

	// defer 是一个栈形式的，例子中，输出 1，4 之后，先输出3，在输出2
	// defer 用于资源的关闭，锁的 lock  defer unlock 会比较方便。
	fmt.Println("---- begin ----")
	for i := 0; i< 10 ;i ++ {
		defer fmt.Println(i)
	}
	fmt.Println("---- finish ----")


	var a,b bool
	for i:=1 ; i<=100 ;i++ {
		a =i%3==0
		b= i%5==0
		if a && b{
			fmt.Println("FizzBizz")
		} else if a {
			fmt.Println("Fizz")
		} else if b{
			fmt.Println("Bizz")
		}
	}

}

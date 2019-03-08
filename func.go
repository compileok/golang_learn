package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func main() {
	fmt.Println(eval(3,4,"*"))
	q,r := divide(10,4)
	fmt.Println(q,r)

	fmt.Println(apply(pow,3,2))

	fmt.Println(sum(1,2,3))

	fmt.Println("---- swap --")
	// swap1
	a ,b := 3,4
	swap1(&a,&b)
	fmt.Println(a,b)
	// swap2
	i,j := 3,4
	i,j = swap2(i,j)
	fmt.Println(i,j)


}

func eval(a,b int, op string) int{
	switch op {
	case "+":
		return a+b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/" :
		return a/b
		/*q ,_ := divide(a,b)
		return q*/
	default:
		panic("unsupported operate: " + op)
	}
}

// 返回两个参数 10 / 4 = 2 余 2
func divide(a ,b int) (q,r int){
	return a / b , a % b
}

func apply(op func(int,int) int , a,b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name();
	fmt.Printf(" calling function : %s, with args: %d , %d \n",opName,a,b)
	return op(a,b)
}

func pow(a,b int) int{
	return int(math.Pow(float64(a) , float64(b)))
}

func sum(numbers...int) int{
	s := 0
	for i:= range numbers {
		s += numbers[i]
	}
	return s
}

//
func swap1(a,b *int) {
	*b,*a = *a,*b
}
func swap2(a,b int)(i,j int){
	return b,a
}


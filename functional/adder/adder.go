package main

import (
	"fmt"
	"strings"
)

func adder() func(int) int{
	sum := 0
	return func(v int) int {
		fmt.Printf(" sum:%d,v:%d \t",sum,v)
		sum += v
		return sum
	}
}
func fibonacci() func() int{
	a ,b := -1,1
	return func() int {
		fmt.Printf(" a:%d,b:%d \t",a,b)
		a,b = b, a+ b
		return b
	}
}
func main() {
	a := adder()
	b := fibonacci()
	for i := 0; i < 4; i++ {
		fmt.Printf(" 0 +  ... %d = %d \n",i,a(i))
	}
	fmt.Println("------")
	for i:= 0 ; i < 4 ;i++ {
		fmt.Printf(" 0 +  ... %d = %d \n",i,b())
	}

	str:="Hello world! hello hello world"
	res := strings.Fields(str)
	smap := map[string]int{}

	fmt.Println("-------")
	for _,v := range res{
		if ele,ok := smap[v];ok{
			smap[v] = ele+1
		}else{
			smap[v] = 1
		}
		fmt.Println(v)
	}
	fmt.Println("----\n",smap)
}

package main

import "fmt"

func main() {

	fmt.Println(" ---- create slice ----")
	var s []int
	for i := 0; i < 22; i++ {
		printSlice(s)
		s = append(s, i) // append 会改变len 和 cap , cap的增长是 1，2，4，8，16，32 cap不够的时候扩容是采取 2的n次方
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	printSlice(s2)

	s3 := make([]int, 10, 32)
	printSlice(s3)

	fmt.Println("---- copy slice ----")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("---- delete element form slice ----")
	//s2[:3] + s2[4:]
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("---- pop from front ----")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(" front", front)
	printSlice(s2)
	fmt.Println("---- pop from back ----")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println("tail", tail)
	printSlice(s2)

}

func printSlice(s []int) {
	fmt.Printf("%v, len = %d, cap = %d \n", s, len(s), cap(s))
}

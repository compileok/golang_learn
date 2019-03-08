package main

import "fmt"

func main() {
	arr := [...]int {1,2,3,4,5,6,7,8,9}

	fmt.Println("arr[2,5]=",arr[2:5])
	fmt.Println("arr[:5]=",arr[:5])
	fmt.Println("arr[2:]=",arr[2:])
	fmt.Println("arr[:]=",arr[:])

	// Slice 本身是没有数据的，是对底层 arry 的一个 view

	// 这里执行以下修改，可以运行看结果，体会一下 slice 是view的概念
	s1 := arr[2:5]
	fmt.Println(" s1 : ",s1)
	updateSlice(s1)
	fmt.Println(" after update ...")
	fmt.Println(" arr: ",arr," s1: ",s1)

	// reslice
	fmt.Println(" reslice s1")
	s1 = s1[2:]
	fmt.Println("after reslice: ",s1)

	// extending slice
	fmt.Println(" extending slice")
	arr2 := [...]int{0,1,2,3,4,5,6,7}
	sa := arr2[2:6] // 2,3,4,5
	sb := sa[3:5] // 5,6
	fmt.Println(" sa = ",sa)
	fmt.Println(" sb = ",sb)

	/**
						0				sb  [5,6]
			0	1	2	3				sa	[2,3,4,5]
	0	1	2	3	4	5	6	7
	 */
}

func updateSlice(s []int) {
	s[0] = 100
}


// arr := [...]int{0,1,2,3,4,5,6,7}
// s1 := arr[2,6]		-> [2,3,4,5]
// s2 := s1[3,5]		-> [5,6]
// slice 内部 有三个元素   ptr len cap
// slice 可以向后拓展，不可以向前拓展
// s[i] 不可超越  len(s) , 向后扩展不可以超过底层数组的 cap(s)
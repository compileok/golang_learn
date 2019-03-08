package main

import "fmt"

func main() {
	// 定义数组
	var arr1 [5]int
	arr2 := [3]int{1,2,3}
	arr3 := [...]int{2,4,6,8}
	fmt.Println(arr1,arr2,arr3)

	// 定义二维数组
	var grid [4][5]int
	fmt.Println(grid)

	// 遍历数组 1
	for i:=0 ; i < len(arr3); i++ {
		fmt.Print(arr3[i],"\t")
	}
	fmt.Println()
    // 遍历数组 2
	for i:= range arr3 {
		fmt.Print(arr3[i],"\t")
	}
	fmt.Println()
	for _,v := range arr3{
		fmt.Print(v,"\t")
	}

	printArray(arr1)
	// printArray(arr3)

}


func printArray(arr [5]int) {
	for _,v := range arr{
		fmt.Print(v,"\t")
	}
}


// [10]int 和 [20]int 是不同类型 （尽管都是数组）
// 调用 func f(arr [10]int) 会 拷贝 数组
// 在 go 语言中一般不直接使用数组   （使用切片）
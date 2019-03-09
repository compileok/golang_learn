package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(
		convert2Bin(5),
		convert2Bin(13))

	printFile("abc.txt")
}

func convert2Bin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

// 这就是一个 while循环，for不写条件就是一个while，所以 为啥要多一个while关键字呢
func whileLike() {
	for {
		fmt.Println("hello")
	}
}

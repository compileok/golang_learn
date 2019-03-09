package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "hi大家好"
	fmt.Println(len(s)) // 这里得到的是 11 ，因为使用了utf-8编码，中文3字节，2 + 9 = 11

	// 打印一下 s 的 字节看看
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	//68 69 E5 A4 A7 E5 AE B6 E5 A5 BD
	fmt.Println()
	for i, ch := range s { // ch is a rune
		fmt.Printf(" %d,%X  ", i, ch)
	}
	// 0,68   1,69   2,5927   5,5BB6   8,597D
	fmt.Println()

	fmt.Println(" rune count : ", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("%d,%c ", i, ch)
	}

}

// rune 相当于 go 的 char

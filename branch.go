package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	contents,err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n",contents)
	}

	grade(100)
	grade(98)
	grade(81)
	grade(70)
	grade(50)
	grade(0)
	grade(-1)
	grade(200)
}

// switch 不需要 显式的写 break
func grade(score int) string{
	g := ""
	switch  {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Errorf(" wrong score :%d",score))
	}
	return g
}
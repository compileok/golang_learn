package main

import (
	"fmt"
	"golang_learn/retriever/mock"
	"golang_learn/retriever/real"
)
type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.so.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{" it is a fake impl"}

	fmt.Printf("%T %v \n",r,r)
	//fmt.Println(download(r))

	r = real.Retriever{}
	fmt.Printf("%T %v \n",r,r)
	//fmt.Println(download(r))

	inspect(r)

	// Type assertion
	if mockRetriever,ok := r.(mock.Retriever);ok {
		fmt.Println(mockRetriever.Contents)
	}else {
		fmt.Println(" not a mock retriever")
	}

}

func inspect(r Retriever) {
	switch v:= r.(type) {
	case mock.Retriever:
		fmt.Println(v.Contents)
	case real.Retriever:
		fmt.Println(v.UserAgent)

	}
}

// 接口变量中有什么，有 实现者的类型，实现者的指针。指针指向实现者
// interface{} 表示任意类型  An empty interface may hold values of any type
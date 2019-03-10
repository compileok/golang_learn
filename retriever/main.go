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
	fmt.Println(download(r))

	r = real.Retriever{}
	fmt.Println(download(r))

}

package main

import "fmt"

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.so.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{""}
	fmt.Println(download(r))

}

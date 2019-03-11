package main

import (
	"errors"
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println(" error occurred : ", err)
		} else {
			panic(r)
		}
	}()
	panic(errors.New(" this is recover"))
}

func main() {
	tryRecover()
}

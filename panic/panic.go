package main

func main() {
	// 相当于别的语言里边的 throw ，要尽量少用  恐慌
	// 停止当前函数，一直向上返回，执行每一层的的defer
	// 如果没有遇见 recover，程序退出

	// recover
	// 仅在defer中调用，获取panic 的值，如果无法处理，可以重新 panic。
}

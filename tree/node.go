package tree

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() { // 结构体方法，前面的（node treeNode）叫做接受者，相当于java 中的this
	fmt.Println(node.value)
}
func print(node treeNode) { // 这里很奇怪？
	fmt.Println(node.value)
}
func (node *treeNode) setValue(v int) {
	node.value = v
}

func main() {

	root := treeNode{}
	root.value = 1
	root.left = &treeNode{}
	root.right = &treeNode{3, nil, nil}
	root.right.left = new(treeNode)
	root.left.left = createNode(9)

	root.setValue(11)
	fmt.Println(" set value:", root)
	root.print()

	nodes := []treeNode{
		{value: 5},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)

	print(root)

}

// go 语言没有构造函数这个说法。如果需要自定义的"构造函数"，可以使用工厂函数,如下 createNode

func createNode(v int) *treeNode {
	return &treeNode{value: v}
}

// 那么结构是分配在堆上还是栈上？ 不需要知道。 go 有垃圾回收器，分配在那里又编译器和运行环境决定，

// 值接收者 Vs 指针接收者
// 要改变内容必须使用指针接收者，结构过大也考虑使用指针接收者
// 一致性：如果有指针接收者，最好都是指针接收者

// 值接收是go语言 特有

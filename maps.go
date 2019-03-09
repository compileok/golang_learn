package main

import "fmt"

func main() {

	m := map[string]string{
		"name":    "zhangsan",
		"address": " beijing china",
		"hello":   "world",
	}

	fmt.Println(m)

	m2 := make(map[string]int) // m2 == empty map
	fmt.Println(m2)

	var m3 map[string]int // m3 == nil
	fmt.Println(m3)

	fmt.Println("---- traversing map ----")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("---- get value ----")
	address := m["address"]
	fmt.Println(address)

	if address2, ok := m["ddress"]; ok {
		fmt.Println(address2)
	} else {
		fmt.Println(" current key does not exists ")
	}

	fmt.Println("---- delete value ----")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}

// 创建： make(map[string]int)
// 获取： m[key]
// key不存在时，获得的value类型的初始值
// 用value，ok：=m[key]来判断是否存在key
// 用delete删除一个key
// 遍历， 使用range， 不保证顺序，如需要顺序 那么就手动对key排序
// 使用 len来获取个数
// map 的 key 类型，map使用哈希表 key 必须可以比较相等；除了slice，map，function的内建类型可以作为key

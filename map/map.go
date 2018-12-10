package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1

	n := map[string]float32{
		"zhangsan": 99.5,
		"lisi":     99.5,
	}

	name, ok := n["zhangsan"]

	num := len(n)

	fmt.Println(num)

	fmt.Println(name, ok)

	//x, ok := m["b"]
	//fmt.Println(x, ok)
	//x, ok = m["a"]
	//fmt.Println(x, ok)
}

package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// 使用 * 读取地址上的值
	fmt.Println(*b)
	fmt.Println(*&a)

	// 通过指针更改值
	*b = 10
	fmt.Println(a)
}

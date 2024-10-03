package main

import "fmt"

func main() {
	// 	bool
	// string
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte // uint8 的别名
	// rune // int32 的别名
	// 表示一个 Unicode 码位
	// float32 float64
	// complex64 complex128

	// 使用 var
	// var name = "Colin"
	var age int32 = 32
	const isCool = true
	var size float32 = 2.3

	// 速记
	// name := "Colin"
	// email := "colinting1001@gmail.com"

	name, email := "Colin", "colinting1001@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
}

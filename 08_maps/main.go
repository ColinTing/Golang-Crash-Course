package main

import "fmt"

func main() {
	// 定义map
	// emails := make(map[string]string)

	// 复制map
	// emails["Colin"] = "colinting1001@gmail.com"
	// emails["John"] = "john@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// 声明map并赋值
	emails := map[string]string{"Colin": "colinting@gmail.com", "John": "john@gmail.com"}

	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Colin"])

	// 删除map中的元素
	delete(emails, "John")
	fmt.Println(emails)
}

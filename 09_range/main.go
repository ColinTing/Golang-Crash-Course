package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	// 循环ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// 不使用索引
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// 累加ids
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// 迭代map
	emails := map[string]string{"Colin": "colinting@gmail.com", "John": "john@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}

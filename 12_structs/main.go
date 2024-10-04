package main

import (
	"fmt"
	"strconv"
)

// 定义一个person结构体
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// 打招呼方法 （值接收者）
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday方法 （指针接收者）
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried （指针接收者）
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// 初始化person
	person1 := Person{firstName: "Colin", lastName: "Ting", city: "Shenzhen", gender: "M", age: 31}

	// 另一种初始化方式
	person2 := Person{"Samantha", "Smith", "Shenzhen", "F", 25}
	fmt.Println(person2)

	// 获取单个字段
	fmt.Println(person1.firstName)
	// person1.age++
	person1.hasBirthday()
	// getMarried方法
	person1.getMarried("Ting")
	person2.getMarried("Ting")

	fmt.Println(person1)

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

}

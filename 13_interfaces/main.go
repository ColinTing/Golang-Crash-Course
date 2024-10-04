package main

import "fmt"

// 定义一个接口
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// 定义 圆形 的方法
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// 定义 矩形 的方法
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// 定义一个方法，接收一个 Shape 类型的参数
func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}

package main

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Person struct {
	Age  int
	Name string
}

type employee struct {
	Person
	EmployeeId int
}

func (e employee) Printinfo() {
	fmt.Println("EmployeeID: %s\n", e.EmployeeId)
	fmt.Println("Name: %s\n", e.Name)
	fmt.Println("Age: %s", e.EmployeeId)

}
func main() {
	rect := Rectangle{3, 4}
	fmt.Println("第一题：矩形面积： %s\n", rect.Area())
	fmt.Println("第一题，矩形周长： %s\n", rect.Perimeter())
	circle := Circle{5}
	fmt.Println("第一题：圆形面积： %s\n", circle.Area())
	fmt.Println("第一题：圆形周长： %s\n", circle.Perimeter())

	emp := &employee{
		Person: Person{
			Age:  10,
			Name: "james",
		},
		EmployeeId: 007,
	}
	emp.Printinfo()
}

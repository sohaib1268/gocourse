package intermediate

import "fmt"

type MyInt int

// methods can be associated with any type and not just structs
func (m MyInt) isPositive() bool { //we declare instance (m) to access values inside
	return m > 0
}

func (MyInt) welcomeMessage() string { //we can declare without instance aswell but we cant access values then
	return "Welcome to MyInt type"
}

type Shape struct {
	Rectangle //embedded struct
}

type Rectangle struct {
	width  int
	height int
}

func (r Rectangle) area() int {
	return r.width * r.height
}

func methods_learning() {
	m1 := MyInt(-5)
	fmt.Println(m1.isPositive())
	m2 := MyInt(5)
	fmt.Println(m2.isPositive())

	fmt.Println(m1.welcomeMessage())
	fmt.Println(m2.welcomeMessage())

	shape := Shape{Rectangle{width: 10, height: 20}}
	fmt.Println("Area of Rectangle:", shape.area())           //we can call area function directly since it was embedded
	fmt.Println("Area of Rectangle:", shape.Rectangle.area()) //same output
}

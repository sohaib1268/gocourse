package intermediate

import (
	"fmt"
	"math"
)

// use interfaces for polymorphism and dynamic, reusability and modular & maintainable code
type geometry interface { //if i make the "g" capital, it will be exportable
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type rect1 struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 { //area is the method defined inside the geometry interface thats why interface implemented
	return r.width * r.height
}

func (r rect1) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func interfaces_learning() {
	r := rect{width: 10, height: 20}
	c := circle{radius: 5}
	//r1 := rect1{width: 10, height: 20}
	measure(r) // rect and circle can be substituted for the geometry type
	measure(c)
	//measure(r1) //will throw error as we have not implemented the perim method

	myPrinter(1, "123", true, r, c)
}

func myPrinter(i ...interface{}) { //any number of interfaces //interface in parameter mean i can use any type of parameter
	for _, v := range i {
		fmt.Println(v)
	}
}

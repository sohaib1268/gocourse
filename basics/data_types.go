package basics

import "fmt"

// middlename := "john" // cant use gofer notation globally, only at function level
var middleName = "Cane"

func basics() {
	fmt.Println("Hello World to data types")

	var age int
	var name string = "John"
	var name1 = "Smith"

	count := 20
	name2 := "Ali"

	middleName := "Cane2"
	fmt.Println(middleName)

	fmt.Println(name + name1 + name2)
	fmt.Println(age + count)
	printname()
	fmt.Println(middleName)

}

func printname() {
	name1 := "afridi"
	fmt.Println("name of person in function " + name1)
}

package intermediate

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	person //embedded struct //employee has inherited person //anonymous field
	//employeeInfo person // embedded struct named field
	empID  string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hello, my name is %s and i Am %d years old", p.name, p.age)
}

// overriding functions
func (emp Employee) introduce() {
	fmt.Printf("Hello, my name is %s and employee ID is %s and i Am %d years old", emp.name, emp.empID, emp.age)
}

func struct_embedding_learning() {
	//embedded struct means inheritance
	emp := Employee{
		person: person{
			name: "John",
			age:  30,
		},
		empID:  "E123",
		salary: 5000,
	}
	fmt.Println(emp.name) //only if anonymous field promotion happens
	//fmt.Println(emp.employeeInfo.name) //only if named field
	fmt.Println(emp.age)
	fmt.Println(emp.empID)
	fmt.Println(emp.salary)

	emp.introduce() // introduce was of person struct but employee can access the method now
	emp.introduce() // overriden prints the employee overriden print statement

}

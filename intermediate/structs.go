package intermediate

import "fmt"

type Person struct {
	firstName     string
	lastName      string
	age           int
	address       Address
	PhoneHomeCell //anonymous fields
	// PhoneHomeCell is an embedded struct, it will be promoted to Person struct
	// so we can access home and cell directly from Person struct
	// home and cell will be accessible as p.home and p.cell
	// if we want to access address.city and address.country then we can use p.address.city and p.address.country
}

type Address struct {
	city    string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}

func (p Person) FullName() string { //define the function is associated with the struct Person
	//Seperation of Concerns (SoC) principle
	//Method to return full name of the person
	//value receiver
	return p.firstName + " " + p.lastName
}

func (p *Person) incremenAgeByOne() {
	//Method to increment age by one
	p.age++
	//Note: we are using pointer receiver here to modify the original struct
	//if we use value receiver, then it will create a copy of the struct and modify the copy
	//so the original struct will not be modified
}

func structs_learning() {

	//Structs are composite data types that group together variables (fields) under a single name
	//Structs don't support inheritance
	//In GO we don't have Classes

	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{ //EMBEDDED STRUCT
			city:    "New York",
			country: "USA",
		},
		PhoneHomeCell: PhoneHomeCell{ //EMBEDDED STRUCT
			home: "123-456-7890",
			cell: "987-654-3210",
		},
	}

	p1 := Person{
		firstName: "John",
		age:       30,
	}

	p1.address.city = "Los Angeles" //we can set the address city after initialization
	p1.address.country = "USA"      //we can set the address country after initialization

	fmt.Println(p.firstName)       //John
	fmt.Println(p.lastName)        //Doe
	fmt.Println(p.age)             //30
	fmt.Println(p.address)         //{New York USA}
	fmt.Println(p.FullName())      //John Doe
	fmt.Println(p1.lastName)       // empty string as LastName is not initialized
	fmt.Println(p1.address)        //{Los Angeles USA}
	fmt.Println(p.address.city)    //New York
	fmt.Println(p.address.country) //USA
	fmt.Println(p.home)            //123-456-7890
	fmt.Println(p.cell)            //987-654-3210

	//Anonymous Structs
	anon := struct {
		Name string
		Age  int
	}{
		Name: "Alice",
		Age:  25,
	}

	fmt.Println(anon.Name) //Alice
	fmt.Println(anon.Age)  //25

	fmt.Println("Age Before Increment", p.age) //30
	p.incremenAgeByOne()                       //increment age by one
	fmt.Println("Age After Increment", p.age)  //31

	// You can compare 2 structs of the same type
	fmt.Println(p == p1)                     //false, because they are not equal
	fmt.Println(p.firstName == p1.firstName) //true, because firstName is the same
	// You can also compare structs with embedded structs
	fmt.Println(p.address == p1.address) //false, because address is not the same

}

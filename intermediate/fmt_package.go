package intermediate

import "fmt"

func fmt_package_learning() {

	//Printing Functions
	fmt.Print("Hello")
	fmt.Print("World!")
	fmt.Print(12, 456)

	fmt.Println("Hello")
	fmt.Println("World!")
	fmt.Println(12, 456)

	name := "John"
	age := 25
	fmt.Printf("Name: %s, Age %d", name, age)
	fmt.Printf("Binary: %b, Hex : %X\n", age, age)

	//output
	// HelloWorld!12 456Hello
	// World!
	// 12 456
	// Name: John, Age 25Binary: 11001, Hex : 19

	//Formatting functions
	//sprint can be used for string appending aswell
	s := fmt.Sprint("Hello", "World!", 123, 456)
	fmt.Print(s)

	s = fmt.Sprintln("Hello", "World!", 123, 456)
	fmt.Print(s)
	fmt.Print(s)

	sf := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Print(sf)
	fmt.Print(sf)

	//output
	// HelloWorld!123 456Hello World! 123 456
	// Hello World! 123 456
	// Name: John, Age: 25Name: John, Age: 25

	//Scanning functions
	var input1 string
	var input2 int
	fmt.Print("Enter your name and age:")
	// fmt.Scan(&input1, &input2)           // it will keep asking even if i press enter, it will ask for 2 values
	// fmt.Scanln(&input1, &input2)         // it stops scanning as soon as it encounters new line
	// fmt.Scanf("%s %d", &input1, &input2) //exact format, expects space after string according to my format
	fmt.Printf("Name: %s , Age: %d", input1, input2)

	//Error Formatting
	err := checkAge(15)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	//output
	// Error:  Age 15 is too young to drive.

}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive.", age)
	}
	return nil
}

package basics

import "fmt"

func switchStatements() {

	fmt.Println("Hello from Switch Statement")

	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a Weekday")
	case "Saturday", "Sunday":
		fmt.Println("It's a Weekend")
	default:
		fmt.Println("Invalid Day")
	}

	number := 15

	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 20")
	default:
		fmt.Println("Number is 20 or greater")
	}

	switch {
	case number > 14:
		fmt.Println("Number is greater than 14")
		fallthrough
	case number == 15:
		fmt.Println("Number is exactly 15")
	default:
		fmt.Println("Number is 14 or less")
	}

	checkType(42)
	checkType(3.14)
	checkType("Hello")
	checkType(true)

}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("x is an integer")
		//fallthrough   // cannot fallthrough in type switch
	case float64:
		fmt.Println("x is a float")
	case string:
		fmt.Println("x is a string")
	default:
		fmt.Println("Unknown type")
	}
}

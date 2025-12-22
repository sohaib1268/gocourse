package basics

import (
	"errors"
	"fmt"
)

func multipleReturnValues() {
	fmt.Println("Hello from Multiple Return Values")
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	result, err := compare(5, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Comparison Result:", result)
	}

	result, err = compare(20, 20)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Comparison Result:", result)
	}
}

func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func smartDivide(a int, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func compare(a int, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater") //for strings nil is "" // errors can be nil as they are interfact
	}
}

package basics

import (
	"fmt"
	"math"
)

func arithmeticOperators() {
	fmt.Println("Hello from arithmetic operators")

	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition :", result)

	result = a - b
	fmt.Println("Subtraction :", result)

	result = a * b
	fmt.Println("Multiplication :", result)

	result = a / b
	fmt.Println("Division :", result)

	result = a % b
	fmt.Println("Remainder (Modulus) : ", result)

	const p float64 = 22 / 7.0
	fmt.Println("Division with 1 operand as non integer : ", p)

	//example of overflow with SIGNED Integer
	var maxInt int64 = 9223372036854775807 // max value that int64 can hold
	fmt.Println(maxInt)

	maxInt += 1
	fmt.Println("After overflow : ", maxInt)

	//example of overflow with UNSIGNED Integer
	var uMaxInt uint64 = 18446744073709551615
	fmt.Println(uMaxInt)

	uMaxInt += 1
	fmt.Println("After overflow : ", uMaxInt)

	//example of underflow

	var smallFloat float64 = 1.0e-323
	fmt.Println("Small float : ", smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println("Underflow example : ", smallFloat)

}

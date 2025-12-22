package basics

import "fmt"

func variadic_functions() {
	fmt.Println("Hello from variadic functions")

	fmt.Println("Sum of 1,2,3:", sum(1, 2, 3))

	fmt.Println(anotherSum("Sum", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(anotherSum("Sum"))

	slice := []int{1, 2, 3, 4, 5, 6, 7}
	sequenceNum, sum := sequenceSum(1, slice...)
	fmt.Println("Sequnce #", sequenceNum, "Sum", sum)
}

//variadic parameters should always be last

func sum(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func anotherSum(returnString string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return returnString, total
}

func sequenceSum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}

package basics

import "fmt"

func functions() {
	fmt.Println("Hello from functions learning")
	fmt.Println(add(2, 3))

	func() {
		fmt.Println("Hello from anonymous function")
	}()

	greet := func() {
		fmt.Println("Hello from assigned function")
	}

	greet()

	//functions as first class citizens
	operation := add
	result := operation(5, 7)
	fmt.Println(result)

	//passing function as argument
	sum := applyOperations(10, 20, add)
	fmt.Println("Sum using applyOperations:", sum)

	//returning and using a function from another function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 =", multiplyBy2(6))

}

func add(a, b int) int {
	return a + b
}

func applyOperations(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

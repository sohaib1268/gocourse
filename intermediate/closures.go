package intermediate

import "fmt"

func closure_learning() {
	fmt.Println("Hello from Closures")
	//commonly used in callback functions
	//avoid useruse
	//can keep variables alive longer than usual
	//must be used carefully to avoid race conditions & memory usage
	sequence := adder()
	fmt.Println(sequence()) //sequence is the anonymous function returned by adder //only increments i // we have i only in memory because of closure
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	//output
	// 	Hello from Closures
	// previous value of i: 0
	// added 1 to i
	// 1
	// added 1 to i
	// 2
	// added 1 to i
	// 3
	// added 1 to i
	// 4

	subtracter := func() func(x int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}() //immediately invoked function expression IIFE

	fmt.Println(subtracter(1))
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(1))

	//output
	// 98
	// 97
	// 96
	// 95
	// 94

}

func adder() func() int {
	i := 0 //scoped to the adder function //will execute only when adder function is called //only once
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}

package basics

import "fmt"

func panic_learning() {
	fmt.Println("Hello from Panic learning in GO")
	process(10)
	process(-5)
}

func process(input int) {
	defer fmt.Println("Deffered statement 1") //these will be executed before the panic unwinds the stack
	defer fmt.Println("Deffered statement 2")

	if input < 0 {
		fmt.Println("Before Panic")
		panic("Negative input not allowed")
		//		defer fmt.Println("Deffered statement 3") //unreachable, nothing executed after panic
	}
	fmt.Println("Processing input:", input)
}

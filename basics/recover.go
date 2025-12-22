package basics

import "fmt"

func recovery_go() {
	fmt.Println("Hello from Recover learning in GO")
	self_heal()
	fmt.Println("Program continues after recovery")
}

func self_heal() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("About to panic")
	panic("Something went wrong!")
	fmt.Println("This line will not be executed")
}

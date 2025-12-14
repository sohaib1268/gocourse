package basics

import "fmt"

func forAsWhile() {
	fmt.Println("Hello from for loop as while")

	i := 1
	for i <= 10 {
		fmt.Println("Iteration :", i)
		i++
	}

	sum := 0
	for {
		sum += 10
		fmt.Println("Current Sum", sum)
		if sum >= 100 {
			fmt.Println("Condition Met", sum)
			break
		}
	}

}

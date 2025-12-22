package basics

import "fmt"

func defer_keyword() {
	fmt.Println("Hello from Defer Learning")
	process(37)

}

// deffered statement are evaluated immediately however, executed at last in LIFO order. so be carefull
func process(i int) {
	fmt.Println("Value of i ", i)
	defer fmt.Println("First deffered statement with i", i)
	defer fmt.Println("Second deffered statement with i", i)
	defer fmt.Println("Third deffered statement with i", i)
	i++
	fmt.Println("Value of i (AGAIN)", i)
}

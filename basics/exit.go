package basics

import (
	"fmt"
	"os"
)

func exit_learning() {
	fmt.Println("Hello from Exit learning in GO")

	defer fmt.Println("This will NOT be printed if os.Exit is called")
	os.Exit(1)
	fmt.Println("This line will never be executed")

}

package basics

import "fmt"

func range_keyword() {
	fmt.Println("Hello from Range learning!")

	message := "Hello Word"

	for i, v := range message {
		fmt.Println("Index:", i, "Value:", v)         //will print ascii value
		fmt.Println("Index:", i, "Value:", string(v)) //will print actual character
		fmt.Printf("Index: %d Rune: %c\n", i, v)      //will also print actual character
	}

}

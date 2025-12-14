package basics

import "fmt"

func for_loop() {

	fmt.Println("Hello from For Loop")

	//simple for loop
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	//iterate over collection
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("Index : %v, Value : %v\n", index, value)
	}

	//dont print even numbers, print odd numbers, and if you encounter the number 5 break out of loop
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even number")
			continue //loop continues but rest of lines skipped
		}
		fmt.Println("Odd Number : ", i)
		if i == 5 {
			fmt.Println("Found Number", i)
			break //break out of loop
		}

	}

	//printing stars //ASTERISK LAYOUT
	var rows int = 5
	for i := 1; i <= rows; i++ {
		//inner loop for spaces before stars
		for j := 1; j < rows-i; j++ {
			fmt.Print(" ")
		}
		//inner loop for stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println() //move to next line
	}

	//New format of for loop
	for i := range 10 {
		fmt.Println(10 - i)
	}
	fmt.Println(" We have a Lift Off !")

}

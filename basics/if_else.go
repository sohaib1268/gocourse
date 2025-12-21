package basics

import "fmt"

func if_else() {

	fmt.Println("Hello from if else in GO")
	temperature := 35

	if temperature >= 30 {
		fmt.Println("It's a hot day")
	} else {
		fmt.Println("It's cool outside")
	}

	if temperature >= 30 {
		if temperature <= 35 {
			fmt.Println("its lesser than or equal to 35 today")

		} else if temperature <= 40 {
			fmt.Println("its lesser than 40 today")
		}
	} else {
		fmt.Println("its less than 30 today")
	}

}

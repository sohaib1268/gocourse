package basics

import "fmt"

const GRAVITY = 9.81

func constants() {
	fmt.Println("hello from constants")

	const (
		Monday           = 1
		Tuesday          = 2
		Wednesday int    = 3
		Thursday  string = "4"
	)

	//NO SHORT DECLARTIONS FOR CONST := , CONVINIENT TO USE CONST BLOCKS like above

	fmt.Println("Thursday" + Thursday)
	fmt.Println(Tuesday)
	fmt.Println(GRAVITY)

}

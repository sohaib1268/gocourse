package intermediate

import "fmt"

func pointers_learning() {
	fmt.Println("Hello from Pointers in GO")
	//a pointer holds the memory address of another variable

	var ptr *int                      //pointer to an integer //nil pointer by default
	fmt.Println("Value of ptr:", ptr) //<nil>
	var a int = 38
	ptr = &a

	fmt.Println("Value of a:", a)      //38
	fmt.Println("Value of ptr:", ptr)  //memory address of a
	fmt.Println("Value at ptr:", *ptr) //38 //dereferencing pointer

	//pointer arithmetic is not allowed in Go
	//pointer operations limited to referencing/assignment and dereferencing

	modifyValue(ptr)
	fmt.Println("Value of a after modifyValue:", a) //39

	//unsafe package allows direct memory access but not recommended for general use
	//unsafe.Pointer(&x) // converts the address of 'x' to an unsafe.Pointer

}

func modifyValue(ptr *int) {
	*ptr++
}

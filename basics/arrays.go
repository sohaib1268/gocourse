package basics

import "fmt"

func arrays() {
	fmt.Println("Hello from Arrays")

	fruits := [4]string{"Apple", "Banana", "Cherry", "Date"}

	fmt.Println("Fruits:", fruits)
	fmt.Println("Third fruit:", fruits[2])

	originalArray := [3]int{100, 200, 300}
	copyArray := originalArray

	copyArray[1] = 201

	fmt.Println("Original Array:", originalArray)
	fmt.Println("Copied Array:", copyArray)

	var integerArray [5]int
	fmt.Println("Integer Array with default values:", integerArray)

	for i, v := range originalArray {
		fmt.Println("Index:", i, "Value:", v)
	}

	// _ is a blank identifier to ignore any value returned that we dont want to use
	for _, v := range fruits {
		fmt.Println("Fruit:", v)
	}

	a, b := someFunction()
	fmt.Println("Values from someFunction:", a, b)

	a, _ = someFunction()
	fmt.Println("Value a from someFunction:", a)

	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("2D Array (Matrix):", matrix)

	anotherOriginalArray := [3]int{10, 20, 30}
	var anotherCopyArray *[3]int

	anotherCopyArray = &anotherOriginalArray
	anotherCopyArray[0] = 11

	fmt.Println("Another Original Array after modification through pointer:", anotherOriginalArray)
	fmt.Println("Another Copy Array (Pointer):", *anotherCopyArray)

}

func someFunction() (int, int) {
	return 3, 8
}

package basics

import (
	"fmt"
	"slices"
)

func slicesLearning() {
	fmt.Println("Hello from Slices")

	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Numbers Slice:", numbers)

	slices1 := [5]int{1, 2, 3, 4, 5}
	slices2 := make([]int, 5)

	a := slices1[1:4]
	fmt.Println("Slice a (slices[1:4]):", a)

	b := slices2[:5]
	fmt.Println("Slice b (slices2[:5]):", b)

	slices2 = append(slices2, 6, 7, 8)
	fmt.Println("Slices2 after append:", slices2)

	slices3 := make([]int, 6)
	copy(slices3, slices2)
	fmt.Println("Slices3 after copy from slices:", slices3)

	if slices.Equal(slices2, slices3) {
		fmt.Println("slices2 and slices3 are equal")
	} else {
		fmt.Println("slices2 and slices3 are not equal")
	}

	fmt.Println("Capacity of slices2:", cap(slices2))
	fmt.Println("Length of slices2:", len(slices2))

}

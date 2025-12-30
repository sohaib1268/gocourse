package intermediate

import "fmt"

func recursion_learning() {
	fmt.Println("Hello from Recursion")
	//Divide and conquer
	//breaks down into smaller subproblems
	//base case to stop recursion else stack overflow
	//recursive case to keep breaking down the problem
	//can be optimized using memoization
	//caching results of expensive function calls and returning the cached result when the same inputs occur again

	fmt.Println(factorial(5))  // 5 * 4 * 3 * 2 * 1 = 120
	fmt.Println(factorial(10)) // 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1 = 3628800
	fmt.Println(factorial(0))  // 1

}

func factorial(n int) int {
	//factorial of 0 is 1
	if n == 0 {
		return 1 //base case
	}
	return n * factorial(n-1) //recursive case
	// n * (n-1) * (n-2) * factorial(n-3)..... factorial(0)
}

package intermediate

import "fmt"

//any is an ALIAS for INTERFACE
//use [ ] brackets for generics
// generics inforce type constraints & reduce type conversion
// generics are useful for reusability

func swap[T any](a T, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(e T) {
	s.elements = append(s.elements, e)
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	e := s.elements[len(s.elements)-1]          // to save and return the last element
	s.elements = s.elements[:len(s.elements)-1] // the : is to remove the element actually // : means start new slice from 0 to len-1 since nothing on left side of :
	return e, true
}

func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Print("Stack Elements : ")
	for _, v := range s.elements {
		fmt.Print(v)
	}
	fmt.Println()
}

func generics_learning() {
	x, y := 1, 2
	fmt.Println("Before Swapping", x, y)
	x, y = swap(1, 2)
	fmt.Println("After Swapping", x, y)

	x1, y1 := "John", "Jane"
	fmt.Println("Before Swapping", x1, y1)
	x1, y1 = swap("John", "Jane")
	fmt.Println("After Swapping", x1, y1)

	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	intStack.printAll()
	fmt.Println(intStack.Pop())
	intStack.printAll()
	fmt.Println("Stack is empty : ", intStack.isEmpty())

	stringStack := Stack[string]{}
	stringStack.Push("John")
	stringStack.Push("Jane")
	stringStack.printAll()
	fmt.Println(stringStack.Pop())
	stringStack.printAll()
	fmt.Println("Stack is empty : ", stringStack.isEmpty())
}

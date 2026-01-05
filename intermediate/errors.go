package main

import (
	"errors"
	"fmt"
)

func sqrt(val float64) (float64, error) {
	if val < 0 {
		return 0, errors.New("cannot compute square root of negative number")
	}
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("data is empty")
	}
	return nil
}

func main() {
	result, err := sqrt(-16)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root is:", result)
	}

	result, err = sqrt(25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root is:", result)
	}

	// data := []byte{}
	// if err := process(data); err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Data processed successfully")

	err = eprocess()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = readData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read successfully")

}

type myError struct {
	msg string
}

func (e *myError) Error() string { //the Error function is built-in function of error interface // We are implementing the Error function of error interface here
	return fmt.Sprintf("Error: %s", e.msg)
}

func eprocess() error {
	return &myError{"Custom error occurred"}
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData failed: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("failed to read config file")
}

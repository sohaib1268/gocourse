package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func strings_runes_learning() {
	fmt.Println("Hello from Strings and Runes")

	//strings are immutable in Go
	// ` ` backticks for raw string literals //will print everything as is including newlines and special characters
	message := "Hello, \nWorld!"
	message2 := "Hello, \tWorld!"
	message3 := "Hello, \rWorld!"

	rawMessage := `Hello, \nWorld!`
	fmt.Println("message", message)
	fmt.Println("message2", message2)
	fmt.Println("message3", message3)
	fmt.Println("raw message", rawMessage)

	//output
	// Hello from Strings and Runes
	// Hello,
	// World!
	// Hello, \nWorld!

	//strings are made up of runes
	fmt.Println("Length of message:", len(message))       //\n is also counted as 1 character
	fmt.Println("Length of rawMessage:", len(rawMessage)) //\n is counted as 2 characters \ and n

	fmt.Println("First character of message:", message[0]) //prints ASCII value of H which is 72

	greeting := "Hello"
	msg := "Ali"
	fmt.Println(greeting + msg) //string concatenation // no spaces added

	//lexical graphical comparison is used to compare strings in Go
	str1 := "Apple"  // A has ASCII value 65
	str2 := "banana" // b has ASCII value 98
	str3 := "app"    //a has ASCII value 97 //compared character by character ascii values
	str4 := "apple"  // a has ASCII value 97

	fmt.Println(str1 < str2) //true
	fmt.Println(str3 < str1) //false
	fmt.Println(str1 < str4) //true
	fmt.Println(str3 < str4) //true

	for i, char := range message {
		fmt.Printf("Character at index %d is %c\n", i, char) //%c for character representation
		fmt.Printf("%x \n", char)                            //hexadecimal representation of rune
		fmt.Printf("%v \n", char)                            //ascii representation of rune
	}

	//output
	// Character at index 0 is H
	// 48
	// 72

	fmt.Println("Rune count :", utf8.RuneCountInString(greeting)) //5

	// we cannot append to original string as strings are immutable in Go
	//but we can create a new string by appending to the original string
	newMessage := greeting + ", " + msg + "!"
	fmt.Println("New Message:", newMessage) //Hello, Ali!

	//Rune is an alias for int32 //represents a Unicode code point
	var r rune = '世'          //single quotes for rune literals
	fmt.Println("Rune r:", r) //19990 //unicode code point for 世
	fmt.Printf("Character for rune r: %c\n", r)

	//output
	// Rune r: 19990
	// Character for rune r: 世

	cstr := string(r)
	fmt.Println("String from rune r:", cstr) //世
	fmt.Printf("Type of cstr: %T\n", cstr)   //string

	const NIHONGO = "日本語"
	fmt.Println(NIHONGO) //日本語

	for _, runevalue := range NIHONGO {
		fmt.Println(runevalue)                   //prints unicode code points for each character
		fmt.Printf("Character: %c\n", runevalue) //prints each character
		//output
		// 26085
		// Character: 日
		// 26412
		// Character: 本
		// 35486
		// Character: 語
	}

	//smiley
	smiley := '😃'
	fmt.Printf("%v\n", smiley)
	fmt.Printf("%c\n", smiley)

	//runes vs char // go vs c
	// runes occupy 4 bytes in memory but in c char occupy 1
	// runes is alias for int32, better for multilingual
	// go has built in unicode support, c traditionally relies on libraries and manual handling along with custom implementation

}

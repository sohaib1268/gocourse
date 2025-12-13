package basics

import (
	"fmt"
	foo "net/http"
)

func imports() {
	fmt.Println("intro to go standard library")

	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP RESPONSE CODE : ", resp.Status)

}

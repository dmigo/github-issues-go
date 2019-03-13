package main

import "net/http"
import "fmt"

func main() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("result: %v\n", resp)
	}

	fmt.Println("termingating")
}

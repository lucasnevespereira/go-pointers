package main

import "fmt"

func main() {
	myValue := 7
	adressOfMyValue := &myValue

	fmt.Println(myValue, adressOfMyValue)

	*adressOfMyValue = 8

	fmt.Println(myValue, adressOfMyValue)
}

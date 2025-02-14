package main

import (
"fmt"
"reflect"
)

func main() {
	var myName string = "Juliana"
	var myAge int = 29
	fmt.Printf("O tipo da variável %v é %T \n", myName, myName)
	fmt.Printf("O tipo da variável %v é %v \n", myAge, reflect.TypeOf(myAge))
}
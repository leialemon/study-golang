package main

import (
	"fmt"
)

func main(){

	frutas := [3]string{"apple", "banana", "guava"}
	fmt.Printf("I have %v fruits! \n", len(frutas))
	
	for index, element := range frutas{
		fmt.Printf("I have %v - %v \n", index, element)
	}

	// Array 2d
	arr := [3][2]int{{2,4},{3,1},{0,2}}
	number := arr[2][0] // == 0

	fmt.Println(number)
}
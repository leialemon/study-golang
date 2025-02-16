package main

import (
	"fmt"
)

func main(){

	var a bool = (3 < 4)
	var b bool = (4 <= 2)
	var c bool = (11 == 11)
	var d bool = (22 != 22)

	fmt.Printf("Is 3 < 4 ? %v \n", a)
	fmt.Printf("Is 4 <= 2 ? %v \n", b)
	fmt.Printf("Is 11 == 11 ? %v \n", c)
	fmt.Printf("Is 22 != 22 ? %v \n", d)

	var one int = 1
	var two int = 2
	three := one + two

	// Se usarmos o operador de adição em strings, elas serão concatenadas
	var abra string = "abra"
	var cadabra string = "cadabra"
	abracadabra := abra + cadabra

	fmt.Println("Arithmetic operators")
	fmt.Println(three)
	fmt.Println(abracadabra)
}
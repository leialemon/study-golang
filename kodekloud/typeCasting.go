package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main(){
	var myInteger int = 35
	var myConvertedFloat float64 = float64(myInteger)
	var myFloat float64 = 3.33
	var myConvertedInteger int = int(myFloat)

	var myAge int = 29
	myConvertedAge := strconv.Itoa(myAge)

	var stringedNumber string = "523"
	numberedNumber, err := strconv.Atoi(stringedNumber)

	var stringedString string = "NaN"
	x,err2 := strconv.Atoi(stringedString)

	fmt.Printf("Variável: myConvertedFloat, Valor: %v, Tipo: %v \n", myConvertedFloat, reflect.TypeOf(myConvertedFloat))
	fmt.Printf("Variável: myConvertedInteger, Valor: %v, Tipo: %v \n", myConvertedInteger, reflect.TypeOf(myConvertedInteger))
	fmt.Printf("Variável: myConvertedAge, Valor: %v, Tipo: %v \n", myConvertedAge, reflect.TypeOf(myConvertedAge))
	fmt.Printf("Variável: numberedNumber, Valor: %v, Tipo: %v, Erro de conversão: %v, Tipo de erro: %v \n", numberedNumber, reflect.TypeOf(numberedNumber), err, reflect.TypeOf(err))
	fmt.Printf("Variável: x, Valor: %v, Tipo: %v, Erro de conversão: %v, Tipo de erro: %v \n", x, reflect.TypeOf(x), err2, reflect.TypeOf(err2))

}
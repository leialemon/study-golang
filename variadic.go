package main

import (
    "fmt")

func soma(numbers ...int) int{
    sum := 0
    for _, element := range numbers{
        sum += element
    } 
    return sum
}

func main(){

    var times int
    fmt.Println("Quantos números você quer somar?")
    fmt.Scanf("%d", &times)

    var numeros [...] int
    for i:=0; i<times; i++{
        var input int
        fmt.Println("Digite o próximo número:")
        fmt.Scanf("%d", &input)
        numeros[i] = input
    }

    fmt.Println(soma(numeros...))

    // Não funciona - não é possível passar a array como argumento da função variádica 
}


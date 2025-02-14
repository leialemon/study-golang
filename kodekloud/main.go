package main // Todo programa Go deve começar com uma declaração do pacote

import "fmt" // Pacote "Format"

func showZeroes() {
    var myString string
    var myInteger int
    var myFloat float64
    var myOtherFloat float32
    var myBoolean bool

    fmt.Println(myString)
    fmt.Println(myInteger)
    fmt.Println(myFloat)
    fmt.Println(myOtherFloat)
    fmt.Println(myBoolean)
}

func main() {
    fmt.Println("Hello World!")
    fmt.Println()
    showZeroes()
}
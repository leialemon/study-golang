# Golang

https://go.dev/learn/

https://go.dev/doc/

Linguagem criada para combinar a facilidade de uma linguagem interpretada e dinamicamente tipada como python e a eficiência e segurança de uma linguagem estaticamente tipada e compilada, como java e c++. Também visava ser moderna, com suporte à computação multinúcleo e em redes.

GO é uma linguagem compilada.

GO é uma linguagem estaticamente tipada, mas a declaração de tipos não precisa ser feita de forma explícita, podendo ser implicitamente deduzida pelo compilador.

## Data Types

- Strings - string;
- Numbers (Integers and Floats) - int, int32, float32, float64;
- Booleans - bool;
- Arrays and Slices;
- Maps;

### Zero values

Quando variáveis são criadas sem nenhum valor atribuído, um valor automático, conhecido como zero value, é alocado a elas. 

O zero value depende do tipo da variável:

- bool - false;
- int - 0;
- float64 - 0.00;
- string “”;
- pointers, functions, interfaces, maps - nil;

## String Formatting

![image.png](Golang%20199dcb0daee08055beeacd07208d2be3/image.png)

## Variáveis

var <nome da variável> <tipo do dado> = <valor da variável>

OU

<nome da variável> := <valor da variável>

```go
var myName string = "Juliana"
var myInteger int = 100
var myBoolean bool = false
var myFloat float64 = 77.90 

// Shorthand
myName := "Juliana"
```

O tipo da variável pode ser obtido usando o especificador de formato %T ou a função TypeOf do pacote reflect:

```go
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
```

Constantes são declaradas com a keyword const:

const <nome da variável> <tipo da variável> = <valor da variável>

```go
const myConstant string = "Uepaa!"
const name = "Nincompoop"
```

Constantes podem ser declaradas com ou sem tipo e precisam ter o valor atribuído já no momento da inicialização.

## Type Casting

É possível converter o tipo de uma variável em GO, mas não há garantias de que o valor dessa variável irá permanecer inalterado.

```go
var myInteger int = 35
var myConvertedFloat float64 = float64(myInteger)
var myFloat float64 = 3.33
var myConvertedInteger int = int(myFloat)
```

### strconv package

Pacote com funcionalidades para converter de e para strings.

- Itoa()
    - Função que converte integers para strings;
    - retorna um valor: a string formada pelo integer convertido.
- Atoi()
    - Função que converte strings em integers;
    - retorna dois valores: o integer convertido e o erro ocorrido no processo (se houver);

## User input

fmt.Scanf(”%format-specifier”, pointer da variável)

```go
package main
import "fmt"

func main(){
	var name string
	fmt.Print("Enter your name:")
	fmt.Scanf("%s", &name)
}
```

A função fmt.Scanf() retorna dois valores: count e err:

- count - o número de argumentos que a função escreve;
- err - algum erro que ocorra durante a execução da função;

## Operadores

Go tem suporte aos operadores de comparação e aritméticos mais comuns.

GO também tem suporte aos operadores lógicos: AND &&, OR ||, NOT !.

Há também os operadores de atribuição:

```go
= //assign
var x int = 3

+= //add and assign
x += 1 // x = 4

-= //subtract and assign
x -= 1 // x = 3

*= //multiply and assign
x *= 2 // x = 6

/= //divide and assign
x /= 2 // x = 3

%= //divide and assign modulus
x %= 2 // x = 1
```

E os operadores bitwise:

& - bitwise AND

> -  bitwise OR
> 

^ - bitwise XOR

>> - right shift

<< - left shift

## Condicionais

if (condição) {

} // os parênteses da condição são opcionais

```go
var age int = 13
if (age >= 18) {
	fmt.Println("Obrigatório votar")
} else if (age < 18 && age >= 16) {
	fmt.Println("Voto facultativo")
} else {
	fmt.Println("Proibido votar")
}
```

switch variável {

case valor_1:

comando

case valor_2:

comando

default:

comando

}

```go
var fruta string = "melancia"
switch fruta{
	case "goiaba":
		fmt.Println("essa é minha fruta favorita!")
	case "manga":
		fmt.Println("eu gosto muito de manga")
	case "maçã":
		fmt.Println("eu não gosto muito de maçã")
	default:
		fmt.Println("eu não conheço essa fruta")
}
```

Switch-case tem uma keyword específica chamada “fallthrough” que força o fluxo da execução para o próximo bloco de case.

Switch-case também pode ser usado com condições, em que cada case tenha ao invés de um valor uma condição a ser avaliada. Quando o switch-case é usado com condições, não se coloca expressões após o “switch”.

Em GO, o “break” do switch-case é implícito, então ele para de ser executado assim que um case é correspondido, a menos que haja fallthrough.

## Loops

for initialization; condition; post {

}

initialization e post são opcionais.

```go
for i := 1; i <= 3; i++ {
	fmt.Println("Hello World")
}

// Sem initialization e post

i := 1 
for i <= 5 {
	fmt.Println(i * i)
	i++
}
```

break sai do loop e continue pula a iteração atual do loop

## Arrays

Coleção de elementos do mesmo tipo que estão armazenados em endereços de memória adjacentes.

O tamanho das arrays em GO é fixo.

As arrays possuem duas propriedades:

- length: referente ao número de elementos na array;
- capacity: referente ao número de elementos que a array pode conter;

var <nome da array> [<tamanho da array>] <tipo do dado>

var <nome da array> [<tamanho da array>] <tipo do dado> = [<tamanho da array>]<tipo do dado>{ elementos da array}

<nome da array> := [<tamanho da array>]<tipo do dado>{elementos da array}

<nome da array> := […]<tipo do dado>{elementos da array}

```go
var frutas [3] string
var notas [2] float64
var frutas2 [3] string = [3]string{"maçã", "banana", "laranja"}
var notas2 [2] float64 = [2]float64{9.2, 7.5}
frutas3 := [3]string{"goiaba", "manga", "maracujá"}

notas3 := [...]float64{8.4, 9,7} 
// Array inicializada com elipses - sem declaração de
//tamanho. O compilador define o tamanho da array pela quantidade de elementos que 
// estão nela.
```

A função len() recebe uma array como parâmetro e retorna o tamanho da array.

É possível acessar os elementos de uma array através de seus índices:

```go
frutas3 := [3]string{"goiaba", "manga", "maracujá"}
novaFruta := frutas3[1] // == "manga"
```

É possível iterar sobre os elementos de uma Array de duas maneiras: usando um for loop comum com o tamanho da array como limite, ou um for loop específico:

for i := 0; i < len(array); i++ {

array[i]

}

for index, element := range array {

}

As arrays também podem ser multidimensionais:

```go
// Array 2d
arr := [3][2]int{{2,4},{3,1},{0,2}}
number := arr[2][0] // == 0
```

## Slices

https://www.w3schools.com/go/go_slices.php

Segmento contínuo de uma array.

Tem 3 principais características:

Pointer - aponta para o primeiro elemento da array que é acessível através daquele slice;

Length - quantidade de elementos que ele contém; len()

Capacity - quantidade de elementos que a array subjacente contém, contados a partir do primeiro elemento do slice. cap()

<slice name> := []<data type>{values}

A partir de uma array: <slice name> := array[index : index]  //o primeiro index é incluído, o segundo não.

Também é possível criar um slice a partir de outro slice com a mesma sintaxe.

<slice name> := make([]<data type>, length, capacity) //cria uma array e um slice.

É possível apendar elementos a um slice com a função append.

## Maps

Coleção não ordenada de pares key/value.

Possuem operações eficientes de add, get e delete.

Tentar colocar pares em mapas não inicializados gerará um erro de runtime.

<nome do mapa> := map[<data type da key>]<data type do value>{pares de key:value}

<nome do mapa> := make(map[<data type da key>]<data type do value>,<capacidade inicial>)

```go
reservas := map[int]string{101:"Smith", 102:"Jones"}
room := reservas[101] // == "Smith"
```

Ao acessar um elemento de um mapa com uma key, dois valores são retornados: o valor correspondente à key e um boolean que informa se aquele key/value foi encontrado no mapa. Se não foi encontrado, o valor retornado será o zero default do data type.

delete(mapa, chave) // exclui o par key:value do mapa.

## Funções

func <nome da função> (parâmetros) <retorno> {

// corpo da função

}

```go
func calcularAreaQuadrado(lado float64) float64{
	return lado * lado
}
```

### Tipos de retorno de funções

É possível retornar múltiplos valores:

```go
func operation(a int, b int) (int, int){
	sum := a + b
	diff := a - b
	return sum, diff
}

func moreOperations(a int, b int) (times int, divided int){
	times = a * b  // Não é preciso inicializar as variáveis, pois elas já foram 
	divided = a / b  // inicializadas na assinatura da função
	return  // Não é preciso especificar quais parâmetros serão retornados
}
```

Funções variádicas - aceitam uma quantidade variável de argumentos do mesmo tipo referenciado na assinatura da função. A função fmt.Println() é uma função variádica.

Nesse caso, os argumentos são armazenados num slice.

func <nome da função> (argumento 1 tipo, argumento 2 …tipo) <tipo do retorno>

```go
func soma(numbers ...int) int{
	sum := 0
	for _, element := range numbers{
		sum += element
	} 
	return sum
}
```

### Funções anônimas

São declaradas sem identificadores, mas podem aceitar inputs e devolver outputs.

Exemplo:

```go
x := func(1 int, b int) int {
	return 1 * b
}
```

### High Order functions

São funções que recebem outras funções como parâmetros ou retornam outras funções como output.

### Defer

Atrasa a execução de uma função até que a função externa termine sua execução.

## Pointers

& (address of) - operador que diz de onde é o endereço que o pointer está guardando;

* (dereference) - operador que retorna o valor que está guardado no endereço;

x := 77 / guardada na memória em 0x0301

&x = 0x0301

*0x0301 = 77

Pointers são declarados com um asterisco, que é diferente do operador deference:

var <nome do pointer> *<tipo do dado> = &<variável>

var <nome do pointer> = &<variável>

<nome do pointer> := &<variável>

## Structs

Data types definidos pelo usuário.

Podem agrupar diferentes variáveis e diferentes data types.

type <nome do struct> struct{

}

```go
type circle struct{
	x float64
	y float64
	z float64
}

type student struct{
	name string
	grades float64[]
}
```

Com o especificador de formato %+v é possível ver todos os valores de todos os atributos de um struct.

Um struct pode ser inicializado de diversas maneiras:

var <nome da  variável> <Struct>

<nome da variável> := new (<Struct>)

<nome da variável> := <Struct> {

<atributo>: <valor>,

<atributo>:<valor>

}

Os campos de um struct podem ser acessados usando .

<nome da variável>.<nome do campo>

```go
type estudante struct{
	nome string
	idade int
}
aluno1 estudante{
	nome:"Maria",
	idade:15
}
fmt.Println(aluno1.idade) // == 15
```

Para atribuir métodos a um struct, basta criar funções (fora do escopo do struct) com referência ao objeto:

```go
type Rectangle struct {
        length  int
        breadth int
}

func (r Rectangle) area() int {
        return r.length * r.breadth
}
```

## Interfaces

type <nome da interface> interface{

//assinatura dos métodos (não pode haver corpo nem variáveis nos métodos)

}

Não há uma maneira explícita de implementar interfaces. A maneira de implementá-las é usando todos os seus métodos.

## FMT Package

print() - função print comum;

println() - função print com um \n ao final;

printf() - função print com suporte a formatação de strings;

scanf() - input do usuário;

## Comandos

- go version - mostra a versão da linguagem instalada;
- go help - lista os comandos disponíveis;
- go run {file.go} - executa um programa .go;
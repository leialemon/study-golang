package main
import "fmt"

func main(){
	var name string
	fmt.Print("Enter your name:")
	fmt.Scanf("%s", &name)
	fmt.Printf("Olá, %s! \n", name)
}
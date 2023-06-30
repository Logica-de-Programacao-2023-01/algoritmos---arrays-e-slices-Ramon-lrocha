package main

import "fmt"

func main() {
	var x string
	fmt.Print("digite uma frase")
	fmt.Scan(&x)
	s := x
	fmt.Println("len(s) =", len(s))

}

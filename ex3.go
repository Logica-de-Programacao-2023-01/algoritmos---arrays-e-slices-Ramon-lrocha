package main

import "fmt"

func main() {
	str := "banana"
	char := 'a'
	count := 0

	for _, c := range str {
		if c == char {
			count++
		}
	}
	fmt.Printf("o caracter %c aparece %d vezes na string %s\n", char, count, str)

}

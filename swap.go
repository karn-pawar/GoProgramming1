package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var t int = a
	a = b
	b = t
	fmt.Println("swapping two numbers:", a, b)
}

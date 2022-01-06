package main

import "fmt"

func main() {
	var a = 4   //int
	var b = 5.3 //float64
	b = float64(a)
	//a = int(b)
	fmt.Println(b)
	fmt.Println(a)

}

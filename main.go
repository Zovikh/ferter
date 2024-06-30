package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println()
	fmt.Scanln(&a)
	fmt.Println()
	fmt.Scanln(&b)
	s := a * b
	fmt.Println(s)

}

package main

import "fmt"

var m = 100

func foo() (int, string) {
	return 10, "melo"
}

func main() {
	n := 10
	m := 1000

	fmt.Println(m, n)

	x, _ := foo()
	_, y := foo()

	fmt.Println("x=", x)
	fmt.Println("y=", y)
}

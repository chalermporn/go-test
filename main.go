package main

import (
	"fmt"

	"github.com/chalermporn/go-test/recursive"
)

func main() {
	fmt.Println("Hello")

	fmt.Println("Fac: ", recursive.Fac(3))
	fmt.Println("Pow: ", recursive.Pow(2, 3))
	fmt.Println("Pow1: ", recursive.Pow1(2, 3))
}

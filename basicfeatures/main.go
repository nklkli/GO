package main

import (
	"fmt"
	// "math/rand"
)

const (
	Watersports = iota
	Soccer
	Chess
)

func main() {
	// fmt.Println(rand.Int())
	const price float32 = 275.00
	const tax float32 = 27.50
	const quantity = 2
	fmt.Println("Total:", Chess*(price+tax))
}

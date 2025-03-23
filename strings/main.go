package main

import (
	"fmt"
)

func main() {
	s := "saы"
	fmt.Println("len=", len([]rune(s)))

	fmt.Printf("%8T %[1]v\n", s)
	fmt.Printf("%8T %[1]X %[1]d\n", []rune(s))
	fmt.Printf("%8T %[1]X %[1]d", []byte(s))
}

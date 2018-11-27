package main

import (
	"fmt"
)

func main() {
	var a = [1]int8{
		10,
	}

	var b = [1]*int8{
		&a[0],
	}

	fmt.Println("Value of a =", a[0])
	fmt.Println("Memory address of a =", &a[0])
	fmt.Println("Value of b =", *b[0])
	fmt.Println("Memory address of b =", b[0])
}

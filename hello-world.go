package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello everyone")
	a := bye()
	_x := 1
	fmt.Println("_x:", _x)
	fmt.Println(a)

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func bye() (x int) {
	n, err := fmt.Println("Bye!!!")
	fmt.Println("n:", n, "err:", err)
	return 1 + 1
}

//

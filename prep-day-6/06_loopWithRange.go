package main

import (
	"fmt"
)

func main() {
	var first, second int
	fmt.Println("Input pertama :")
	fmt.Scan(&first)
	fmt.Println("Input kedua :")
	fmt.Scan(&second)
	for i := first; i <= second; i++ {
		fmt.Println(i)
	}
}

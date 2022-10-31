package main

import (
	"fmt"
	"testing"
)

func TestGanjil(t *testing.T) {
	for i := 1; i <= 100; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}

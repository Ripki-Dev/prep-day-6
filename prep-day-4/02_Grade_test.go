package prep_day_4

import (
	"fmt"
	"testing"
)

func TestGrade(t *testing.T) {
	nilai := 89
	Grade(nilai)
}

func Grade(n int) {
	if n >= 90 {
		fmt.Println("A")
	} else if n >= 80 {
		fmt.Println("B")
	} else if n >= 70 {
		fmt.Println("C")
	} else if n >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}

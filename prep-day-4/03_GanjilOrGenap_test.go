package prep_day_4

import (
	"fmt"
	"testing"
)

func GanjilOrGenap(n int) {
	if n%2 == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}
}

func TestGanjilOrGenap(t *testing.T) {
	n := 1032
	GanjilOrGenap(n)
}

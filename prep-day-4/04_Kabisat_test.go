package prep_day_4

import (
	"fmt"
	"testing"
)

func TestKabisat(t *testing.T) {
	year := 2020
	Kabisat(year)
}

func Kabisat(n int) {
	if n%400 == 0 {
		fmt.Println("Tahun Kabisat")
	} else if n%400 >= 1 && n%100 >= 1 && n%4 == 0 {
		fmt.Println("Tahun Kabisat")
	} else {
		fmt.Println("Bukan Tahun Kabisat")
	}
}

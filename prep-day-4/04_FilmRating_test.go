package prep_day_4

import (
	"fmt"
	"testing"
)

func TestFilmRating(t *testing.T) {
	rate := 32
	FilmRating(rate)
}

func FilmRating(n int) {
	if n >= 21 {
		fmt.Println("Dewasa")
	} else if n >= 13 {
		fmt.Println("Remaja")
	} else if n >= 9 {
		fmt.Println("Bimbingan Orang Tua")
	} else {
		fmt.Println("Semua Usia")
	}

}

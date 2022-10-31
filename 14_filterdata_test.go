package main

import (
	"fmt"
	"testing"
)

type People struct {
	id   int
	name string
	age  int
}

func TestFilterData(t *testing.T) {
	var people = []People{
		{
			id:   1,
			name: "Udin",
			age:  12,
		},
		{
			id:   2,
			name: "Wati",
			age:  51,
		},
		{
			id:   3,
			name: "Budi",
			age:  34,
		},
		{
			id:   4,
			name: "Agus",
			age:  16,
		},
		{
			id:   5,
			name: "Sari",
			age:  19,
		},
		{
			id:   6,
			name: "Ririn",
			age:  21,
		},
	}
	for i := 0; i < len(people); i++ {
		if people[i].age < 21 {
			fmt.Println(people[i])
		}
	}
}

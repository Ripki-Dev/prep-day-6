package main

import (
	"fmt"
	"testing"
)

func TestArrRemove(t *testing.T) {
	fruits := []string{"Jeruk", "Pepaya", "Jambu", "Anggur", "Melon"}
	fmt.Println(RemoveArr("Jambu", fruits))
}
func RemoveArr(word string, s []string) []string {
	var result []string
	for i := 0; i < len(s); i++ {
		if s[i] != word {
			result = append(result, s[i])
		}
	}
	return result
}

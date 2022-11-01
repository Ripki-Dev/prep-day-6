package main

import (
	"fmt"
	"testing"
)

func TestAddtoArr(t *testing.T) {
	stuf := []string{"Meja", "Buku", "Topi", "Baju", "Kayu"}
	fmt.Println("Tambah Celana di akhir")
	fmt.Println(tambahAkhir(stuf))
	fmt.Println("Tambah Handuk diawal")
	fmt.Println(tambahAwal(stuf))
	fmt.Println("Disatukan")
	xyz := tambahAkhir(stuf)
	fmt.Println(tambahAwal(xyz))
}
func tambahAkhir(s []string) []string {
	add := "Celana"
	s = append(s, add)
	return s
}

func tambahAwal(s []string) []string {
	add := "Handuk"
	s = append([]string{add}, s...)
	return s
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var bangunDatar hitunghitungan
	bangunDatar = &ppanjang{panjang: 10, lebar: 5}

	fmt.Println(bangunDatar.keliling())
}

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

type hitunghitungan interface {
	hitung2d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

type ppanjang struct {
	panjang float64
	lebar   float64
}

func (pp *ppanjang) luas() float64 {
	return pp.panjang * pp.lebar
}

func (pp *ppanjang) keliling() float64 {
	return 2 * (pp.panjang + pp.lebar)
}

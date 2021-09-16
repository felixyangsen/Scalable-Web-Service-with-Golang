package main

import "math"

func main() {
	var bangunDatar hitung

	bangunDatar = lingkaran{14}
	bangunDatar.keliling()
}

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jarijari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jarijari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

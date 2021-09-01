package main

import "fmt"

func main() {
	
}

func condition1() {
	//if ,else if, else
	var a, b int

	if a > b {
		fmt.Println("a > b")
	} else if a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a = b")
	}
}

func condition2() {
	// variabel temporary if
	if a := 10; a < 5 {
		fmt.Println("a < 5")
	}
}

func condition3() {
	//nested if
	var a, b int

	if a >= b {
		if a == b {
			fmt.Println("a = b")
		} else {
			fmt.Println("a != b")
		}
	}
}

func condition4() {
	// switch case
	var a int

	switch a {
	case 1:
		fmt.Println("a = 1")
	case 2, 3:
		fmt.Println("a = 2 || 3")
	default:
		fmt.Println("default")
	}
}

func condition5() {
	//switch case fallthrough
	var a int = 1

	switch a {
	case 1:
		fmt.Println("a = 1")
		fallthrough
	case 2:
		fmt.Println("a = 2")
	default:
		fmt.Println("a != 1 && a != 2")
	}
}

func array1() {
	// array
	var city [10]string // deklarasi dengan range
	city[0] = "medan"

	var country = [2]string{"indonesia", "singapore"} // inisialisasi langsung
	var car = [...]string{"toyota", "honda"}          // array tanpa range

	var number = [2][2]int{{1, 2}, {1, 2}} // array multidimensi

	fmt.Println(city, country, car, number)
}

func slice1() {
	// slice
	var city []string // deklarasi
	var country = []string{"indonesia", "singapore"} // inisialisasi langsung
	var countryArr = make([]string, 3) // deklarasi dengan make
	country = append(country, "vietnam") // menambah element dalam slice
	copy(countryArr, country) // mengcopy slice ke slice lain

	fmt.Println(len(country) // mencari jumlah element slice
	fmt.Println(cap(country) // mencari maks element slice
	fmt.Println(country[0:2], city)
}
	
func map1() {
	// map
	var city map[int]string // deklarasi

	var country = map[int]string{}
	country[1] = "indonesia" // inisialisasi

	var number = map[int]string{1: "a", 2: "b"} // inisialisasi langsung
	var value, isExist = number[1] //check key and value
	delete(number, 1)              // delete

	fmt.Println(city, country, number, value, isExist)
}

func looping1() {
	// loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// loop without condition
	for {
		a := 5
		if a == 10 {
			break
		}

		a++
	}

	// loop with argument
	var a = 5
	for a < 10 {
		a++
	}

	// loop with array
	var arr []int
	for index, value := range arr {
		fmt.Println(index, value)
	}

	// nested loop
	for i := 0; i < 5; i++ {
		for j := i; j < 5; i++ {
			fmt.Println(i, j)
		}
	}
}

func mapInGo() {
	fruits := make(map[string]int)
	maxInt := 0
	maxStr := ""

	arrFruits := []string{"jeruk", "apel", "jeruk", "semangka", "sirsak", "jeruk", "semangka", "semangka", "sirsak", "jeruk"}

	for _, v := range arrFruits {
		fruits[v] += 1
	}

	for i, v := range fruits {
		if maxInt < v {
			maxInt = v
			maxStr = i
		}
	}

	fmt.Println(maxInt, maxStr)

	// hitung buah yg terbanyak, buat menjadi 1 function
}

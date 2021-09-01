package main

import (
	"fmt"
	"os"
	"strconv"
)

type User struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func quiz() {
	var friends = []User{
		{
			Nama:      "Triyono",
			Alamat:    "jalan jalan",
			Pekerjaan: "flutter dev",
			Alasan:    "buat suplay aplikasi flutter ku",
		},
		{
			Nama:      "Aditya Rizki Pratama",
			Alamat:    "di rumah",
			Pekerjaan: "kurang tau juga",
			Alasan:    "bisa upgrade skill backend, belajar micro service dan pastinya manambah earning karena fee golang developer sangat menarik dibandingakan dengan programming language yang saya gunakan sebelumnya hehe",
		},
		{
			Nama:      "Arfan",
			Alamat:    "indonesia",
			Pekerjaan: "juga gatau",
			Alasan:    "agar dapat memahami lebih dalam bahasa Go terutama implementasi di bidang backend web",
		},
	}

	argsCount := len(os.Args)
	if argsCount <= 1 {
		fmt.Println("no id inserted")
		return
	}

	input := os.Args[1]
	intValue, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(friends[intValue+1])
	return
}

func main() {
	quiz()
}

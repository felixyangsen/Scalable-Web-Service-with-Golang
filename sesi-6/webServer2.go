// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// const port = ":8080"

// type ResponseData struct {
// 	Name    string
// 	Address string
// 	Age     int
// }

// func main() {
// 	http.HandleFunc("/index", index)

// 	fmt.Println("running at http://localhost", port)
// 	http.ListenAndServe(port, nil)
// }

// func index(w http.ResponseWriter, r *http.Request) {
// 	var response ResponseData
// 	response.Name = "adit"
// 	response.Address = "jepara"
// 	response.Age = 21

// 	fmt.Println(response)
// }


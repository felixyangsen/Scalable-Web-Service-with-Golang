// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// const port = ":8080"

// func main() {

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "halo")
// 	})

// 	http.HandleFunc("/index", index)

// 	fmt.Println("running at http://localhost", port)
// 	http.ListenAndServe(port, nil)
// }

// func index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "apa kabar!")
// }


// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// type RequestParam struct {
// 	Param string `json:"param"`
// }

// type ResponseParam struct {
// 	Status int
// 	Desc   bool
// }

// const port = ":8080"

// func main() {
// 	router := mux.NewRouter()

// 	router.HandleFunc("/is-palindrom", isPalindromHandler).Methods("POST")

// 	// http.Handle("/", router)

// 	fmt.Println("running at http://localhost", port)
// 	http.ListenAndServe(port, router)
// }

// func isPalindromHandler(w http.ResponseWriter, r *http.Request) {
// 	var param RequestParam
// 	var response ResponseParam
// 	c, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer r.Body.Close()

// 	if err := json.Unmarshal(c, &param); err != nil {
// 		panic(err)
// 	}

// 	if res := isPalindrom(param.Param); res {
// 		response = ResponseParam{
// 			Status: 200,
// 			Desc:   true,
// 		}
// 	} else {
// 		response = ResponseParam{
// 			Status: 200,
// 			Desc:   false,
// 		}
// 	}

// 	json.NewEncoder(w).Encode(response)
// }

// func isPalindrom(str string) bool {
// 	for i := 0; i < len(str)/2; i++ {
// 		j := len(str) - 1 - i
// 		if str[i] != str[j] {
// 			return false
// 		}
// 	}
// 	return true
// }

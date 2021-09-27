// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	// post request
// 	postBody, _ := json.Marshal(map[string]string{
// 		"name":  "Arfan",
// 		"email": "arfan@gmail.coom",
// 	})

// 	responseBody := bytes.NewBuffer(postBody)

// 	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	sb := string(body)
// 	fmt.Println(sb)
// }

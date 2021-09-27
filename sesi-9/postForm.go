package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// post form
	var result map[string]interface{}
	formData := url.Values{
		"name":  {"hanif"},
		"email": {"hanif@gmail.com"},
	}

	resp, err := http.PostForm("http://httpbin.org/post", formData)
	if err != nil {
		log.Fatalln(err)
	}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result["form"])
}

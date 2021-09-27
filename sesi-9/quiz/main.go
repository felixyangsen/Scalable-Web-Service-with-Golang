package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Relation struct {
	InfluencedBy []string `json:"influenced-by"`
	Influences   []string `json:"influences"`
}

type Data struct {
	Language       string   `json:"language"`
	Appeared       int      `json:"appeared"`
	Created        []string `json:"created"`
	Functional     bool     `json:"functional"`
	ObjectOriented bool     `json:"object-oriented"`
	Relation       Relation `json:"relation"`
}

const port = ":8080"

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/insert-language", insertLanguageHandler)
	router.HandleFunc("/language", getLanguageHandler)
	router.HandleFunc("/is-palindrom", isPalindromHandler)

	router.NotFoundHandler = http.HandlerFunc(notfoundHandler)

	http.Handle("/", router)

	fmt.Println("running at http://localhost", port)
	http.ListenAndServe(port, nil)
}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("method not allowed"))
}

func insertLanguageHandler(w http.ResponseWriter, r *http.Request) {
	var param Data

	c, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	if err := json.Unmarshal(c, &param); err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(param)
}

func getLanguageHandler(w http.ResponseWriter, r *http.Request) {
	var param = Data{
		Language:       "C",
		Appeared:       1972,
		Created:        []string{"Dennis Ritchie"},
		Functional:     true,
		ObjectOriented: false,
		Relation: Relation{
			InfluencedBy: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"},
			Influences:   []string{"C++", "Objective-C", "C#", "Java", "JavaScript", "PHP", "Go"},
		},
	}

	json.NewEncoder(w).Encode(param)
}

func isPalindromHandler(w http.ResponseWriter, r *http.Request) {
	var response string
	var param = r.URL.Query()["text"]

	if len(param) == 0 {
		response = "not parameter found"
		json.NewEncoder(w).Encode(response)
		return
	}

	if res := isPalindrom(param[0]); res {
		response = "Palindrome"
	} else {
		response = "Not palindrome"
	}

	json.NewEncoder(w).Encode(response)
}

func isPalindrom(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

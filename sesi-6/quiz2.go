// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// type Relation struct {
// 	InfluencedBy []string `json:"influenced-by"`
// 	Influences   []string `json:"influences"`
// }

// type Data struct {
// 	Language       string   `json:"language"`
// 	Appeared       int      `json:"appeared"`
// 	Created        []string `json:"created"`
// 	Functional     bool     `json:"functional"`
// 	ObjectOriented bool     `json:"object-oriented"`
// 	Relation       Relation `json:"relation"`
// }

// const port = ":8080"

// func main() {
// 	router := mux.NewRouter()

// 	router.HandleFunc("/get-data", getDataHandler).Methods("POST")

// 	http.Handle("/", router)

// 	fmt.Println("running at http://localhost", port)
// 	http.ListenAndServe(port, nil)
// }

// func getDataHandler(w http.ResponseWriter, r *http.Request) {
// 	var param Data

// 	c, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer r.Body.Close()

// 	if err := json.Unmarshal(c, &param); err != nil {
// 		panic(err)
// 	}

// 	json.NewEncoder(w).Encode(param)
// }

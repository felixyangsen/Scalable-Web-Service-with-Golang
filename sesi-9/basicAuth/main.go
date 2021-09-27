package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/students", ActionStudent)

	server := new(http.Server)
	server.Addr = ":8081"

	fmt.Println("running at localhost:8081")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}

	if !AllowOnlyGet(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		outputJSON(w, SelectStudent(id))
		return
	}

	outputJSON(w, GetStudents())
}

func outputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux

	mux.HandleFunc("/students", ActionStudent)
	mux.HandleFunc("/teachers", ActionTeacher)

	// implementasi 
	var handler http.Handler = mux
	handler = AuthMiddleware(handler)
	handler = AllowOnlyGetMiddleware(handler)

	server := new(http.Server)
	server.Addr = ":8081"
	server.Handler = handler

	fmt.Println("running at localhost:8081")
	server.ListenAndServe()
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

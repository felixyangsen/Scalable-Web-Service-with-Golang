package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

var port = ":8081"

type Status struct {
	Water int
	Wind  int
}

type StatusValue struct {
	WaterValue  int
	WaterStatus string
	WindValue   int
	WindStatus  string
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/get-status", getStatusHandler)

	fmt.Println("running at localhost" + port)
	http.ListenAndServe(port, router)
}

func getStatusHandler(w http.ResponseWriter, r *http.Request) {
	status := GetStatus()
	statusValue := GetStatusValue(status)

	var t, err = template.ParseFiles("template.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	t.Execute(w, statusValue)
}

func randomIntGen(start, end int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(end-start) + start
}

func GetStatus() Status {
	var status Status

	status.Water = randomIntGen(1, 100)
	time.Sleep(1)
	status.Wind = randomIntGen(1, 100)

	return status
}

func GetStatusValue(status Status) StatusValue {
	var statusValue = StatusValue{
		WaterValue: status.Water,
		WindValue:  status.Wind,
	}

	if status.Water <= 5 {
		statusValue.WaterStatus = "aman"
	} else if status.Water >= 6 && status.Water <= 8 {
		statusValue.WaterStatus = "siaga"
	} else if status.Water > 8 {
		statusValue.WaterStatus = "bahaya"
	}

	if status.Wind <= 6 {
		statusValue.WindStatus = "aman"
	} else if status.Wind >= 7 && status.Water <= 15 {
		statusValue.WindStatus = "siaga"
	} else if status.Wind > 15 {
		statusValue.WindStatus = "bahaya"
	}

	return statusValue
}

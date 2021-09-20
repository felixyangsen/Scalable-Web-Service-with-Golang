package service

import (
	"encoding/json"
	"io/ioutil"
	"myapp/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var newOrder model.Order
	var response model.Response

	c, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	if err := json.Unmarshal(c, &newOrder); err != nil {
		panic(err)
	}

	_, err = CreateOrder(newOrder)
	if err != nil {
		panic(err)
	}

	response = model.Response{
		Status:  true,
		Message: "success",
	}

	json.NewEncoder(w).Encode(response)
}

func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := GetOrderList()
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(resp)
}

func UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order model.Order

	c, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	if err := json.Unmarshal(c, &order); err != nil {
		panic(err)
	}

	_, err = UpdateOrder(order)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(order)
}

func DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var vars = mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		panic("id not found")
	}

	intID, _ := strconv.Atoi(id)
	_, err := DeleteOrder(intID)
	if err != nil {
		panic(err)
	}

	response = model.Response{
		Status:  true,
		Message: "deleted",
	}

	json.NewEncoder(w).Encode(response)
}

package service

import (
	"encoding/json"
	"io/ioutil"
	"myapp/model"
	"net/http"
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

package service

import (
	"myapp/config"
	"myapp/model"
)

func CreateOrder(newOrder model.Order) (*model.Order, error) {
	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Create(&newOrder).Error

	return &newOrder, err
}

func GetOrderList() ([]*model.Order, error) {
	var orders []*model.Order

	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Preload("Items").Find(&orders).Error

	return orders, err
}

func UpdateOrder(order model.Order) (*model.Order, error) {
	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Create(&order).Error

	return &order, err
}
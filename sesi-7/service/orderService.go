package service

import (
	"myapp/config"
	"myapp/model"

	"gorm.io/gorm"
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

	err := db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, err
}

func DeleteOrder(orderID int) (int, error) {
	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Where("order_id = ?", orderID).Delete(&model.Order{}).Error

	return orderID, err
}

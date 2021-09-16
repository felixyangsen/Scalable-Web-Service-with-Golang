package model

import "time"

type Order struct {
	OrderID      int       `json:"order_id" gorm:"primaryKey"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID;references:OrderID"`
}
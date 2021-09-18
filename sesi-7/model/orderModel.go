package model

import "time"

type Order struct {
	OrderID      int       `json:"orderId" gorm:"primaryKey"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID;references:OrderID"`
}

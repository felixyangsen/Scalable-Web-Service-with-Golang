package model

type Item struct {
	ItemID      int    `json:"itemId" gorm:"primaryKey"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"orderId"`
}

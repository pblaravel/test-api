package models

import "time"

type Order struct {
	ID        int
	Email     string
	Status    string
	Sum       float32
	Items     []OrderItem
	CreatedAt time.Time
}

type OrderItem struct {
	ID        int
	OrderId   int
	Count     int
	ShopId    string
	Shop      Shop `gorm:"foreignKey:shop_id;references:uuid"`
	ProductId int
	Product   Product
}
type ItemCount struct {
	Count     int
	ProductId int
}
type CreateOrder struct {
	Email string
	Items []ItemCount
}

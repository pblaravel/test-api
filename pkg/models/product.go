package models

type Product struct {
	ID          int `json:"-" gorm:"primary_key"`
	ShopID      string
	Name        string
	Description string
	Photos      string
	Price       float64
	Color       string
	Category    string
}

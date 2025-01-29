package models

type Shop struct {
	UUID        string `json:"-" gorm:"unique"`
	Wallet      string
	Owner       string
	Name        string
	Logo        string
	Location    string
	Category    string
	SubCategory string
	Description string
	Rules       string
	Status      int
}

type Category struct {
	ID   int    `json:"-" gorm:"primary_key"`
	Name string `json:"name" gorm:"unique"`
}

type SubCategory struct {
	ID         int `json:"-" gorm:"primary_key"`
	Name       string
	CategoryID int
}

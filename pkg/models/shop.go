package models

type Shop struct {
	UUID        string `json:"-" gorm:"unique"`
	Name        string
	Logo        string
	Location    string
	Category    string
	SubCategory string
	Description string
	Rules       string
	Status      int
}

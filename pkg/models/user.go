package models

type User struct {
	ID       uint64 `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

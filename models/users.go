package models

type User struct {
	ID       int
	Username string `gorm:"unique"`
	Password string
	Email    string
	Address  string
}

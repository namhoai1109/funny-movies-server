package model

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`

	Links []*Link `json:"links" gorm:"foreignKey:UserID"`

	Base
}

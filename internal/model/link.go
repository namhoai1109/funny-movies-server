package model

type Link struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Url    string `json:"url" gorm:"default:NULL"`
	UserID int    `json:"user_id" gorm:"default:NULL"`

	Base
}

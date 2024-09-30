package model

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`

	Links []*Link `json:"links" gorm:"foreignKey:UserID"`

	Base
}

type UserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:    u.ID,
		Email: u.Email,
	}
}

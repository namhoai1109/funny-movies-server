package model

type Link struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Url    string `json:"url" gorm:"default:NULL"`
	UserID int    `json:"user_id" gorm:"default:NULL"`

	User *User `json:"user" gorm:"foreignKey:UserID"`

	Base
}

type LinkResponse struct {
	ID   int           `json:"id"`
	Url  string        `json:"url"`
	User *UserResponse `json:"user"`
}

func (l *Link) ToResponse() *LinkResponse {
	userRes := &UserResponse{}

	if l.User != nil {
		userRes = l.User.ToResponse()
	}

	return &LinkResponse{
		ID:   l.ID,
		Url:  l.Url,
		User: userRes,
	}
}

package model

import "time"

type UserModel struct {
	Id int `gorm:"primary_key"`
	Username string
	Password string
	Roles string
	ClientId string
	ClientSecret string
	Scope string
	OpenId string
	GrantType string
	AccessToken string
	RefreshToken string
	TokenType string
	ExpiresIn int
	Email string
	Telephone string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserModel) TableName() string {
	return "user"
}

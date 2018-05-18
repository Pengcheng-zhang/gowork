package model

import "time"

type UserModel struct {
	Id int `gorm:"primary_key"`
	Username string
	Password string
	Roles string
	ClientId string
	ClientSecret string
	ScoreTotal int
	ScoreCurrent int
	ScoreCost int
	OpenId string
	GrantType string
	AccessToken string
	RefreshToken string
	TokenType string
	ExpiresIn int
	Email string
	Telephone string
	Verified string
	HeadImage string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserModel) TableName() string {
	return "yz_user"
}

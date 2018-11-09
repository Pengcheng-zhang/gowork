package model

import (
	"time"
)

type FriendsModel struct {
	Id int
	UserId int
	Title string
	Name string
	Sex int
	BirthDay string
	Height string
	Weight string
	Married string
	Education string
	CurrentCity string
	RegistCity string
	BornCity string
	Profession string
	Parents string
	Brothers string
	IsOnly string
	InCome string
	Interest string
	PlaceOther string
	MarryYears string
	ChildNum string
	RequestBase string
	RequestOther string
	ShowMeSpecial string
	SelfRecommend string
	ImageUrlOne string
	ImageUrlTwo string
	ImageUrlThree string
	ImageUrlFour string
	Contact string
	SourceDest string
	CreatedAt time.Time
}

func (FriendsModel) TableName() string {
	return "yz_friends"
}